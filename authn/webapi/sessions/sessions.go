// brian taylor vann
// taylorvann dot com
//
// Package sessions - meta-interface for whitelist, csrfs, and session tokens

package sessions

import (
	"errors"
	"webapi/interfaces/jwtx"
	"webapi/sessions/constants"
	"webapi/store"
	"webapi/whitelist"
)

type MilliSeconds = int64

type Session struct {
	SessionToken string `json:"session_token"`
}

type CreatePublicJWTParams struct {
	Email    string
	Password string
}

type CreateAccountParams struct {
	Email string
}

type CreateParams = jwtx.Claims

type ReadParams struct {
	SessionToken *string
}

type UpdateParams struct {
	SessionToken *string
}

type ValidateAndRemoveParams = UpdateParams

type RemoveParams = whitelist.RemoveEntryParams

func getLifetimeByAudience(audience string) int64 {
	switch audience {
	case constants.Guest:
		return constants.OneDayAsMS
	case constants.Public:
		return constants.ThreeDaysAsMS
	default:
		return constants.OneDayAsMS
	}
}

func ComposeDocumentSessionParams() *CreateParams {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Guest)

	params := CreateParams{
		Iss: constants.TaylorVannDotCom,
		Sub: constants.Guest,
		Aud: constants.Document,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &params
}

func ComposeGuestSessionParams() *CreateParams {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Guest)

	params := CreateParams{
		Iss: constants.TaylorVannDotCom,
		Sub: constants.Guest,
		Aud: constants.Session,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &params
}

func ComposeResetPasswordSessionParams(p *CreateAccountParams) *CreateParams {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Guest)

	params := CreateParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.ResetPassword,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &params
}

func ComposeAccountCreationSessionParams(p *CreateAccountParams) (*CreateParams, error) {
	if p == nil {
		return nil, errors.New("ComposeAccountCreationSessionParams - nil parameters provided")
	}

	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Guest)

	params := CreateParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.CreateAccount,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &params, nil
}

func ComposePublicSessionParams(p *CreatePublicJWTParams) (*CreateParams, error) {
	if p == nil {
		return nil, errors.New("ComposePublicSessionParams - nil parameters provided")
	}
	userRow, errValidUser := store.ValidateUser(
		&store.ValidateUserParams{
			Email:    p.Email,
			Password: p.Password,
		},
	)
	if errValidUser != nil {
		return nil, errValidUser
	}
	if userRow == nil {
		return nil, errors.New("bad credentials provided")
	}

	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Public)

	params := CreateParams{
		Iss: constants.TaylorVannDotCom,
		Sub: string(userRow.ID),
		Aud: constants.Public,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &params, nil
}

func Create(p *CreateParams) (*Session, error) {
	if p == nil {
		return nil, errors.New("nil CreateParams provided")
	}

	token, errToken := jwtx.CreateJWT(p)
	if errToken != nil {
		return nil, errToken
	}

	_, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			CreatedAt:  p.Iat,
			Lifetime:   getLifetimeByAudience(p.Aud),
			SessionKey: token.RandomSecret,
			Signature:  &token.Token.Signature,
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	sessionTokenAsStr, errSessionTokenAsStr := jwtx.ConvertTokenToString(
		token.Token,
	)
	if errSessionTokenAsStr != nil {
		return nil, errSessionTokenAsStr
	}

	session := Session{
		SessionToken: *sessionTokenAsStr,
	}

	return &session, nil
}

func Read(p *ReadParams) (bool, error) {
	if p == nil {
		return false, errors.New("nil params")
	}

	if p.SessionToken == nil {
		return false, errors.New("nil session token provided")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.SessionToken,
	)
	if errTokenDetails != nil {
		return false, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Signature: &tokenDetails.Signature,
		},
	)
	if errEntry != nil {
		return false, errEntry
	}

	if entry != nil {
		result := jwtx.ValidateJWT(&jwtx.TokenPayload{
			Token:        tokenDetails,
			RandomSecret: &entry.SessionKey,
		})
		return result, nil
	}

	return false, nil
}

func ValidateAndRemove(p *ValidateAndRemoveParams) (*whitelist.Entry, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	if p.SessionToken == nil {
		return nil, errors.New("nil sesion provided")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.SessionToken,
	)
	if errTokenDetails != nil {
		return nil, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Signature: &tokenDetails.Signature,
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	if entry != nil {
		resultJwt := jwtx.ValidateJWT(&jwtx.TokenPayload{
			Token:        tokenDetails,
			RandomSecret: &entry.SessionKey,
		})
		if resultJwt {
			removeResult, errRemoveResult := whitelist.RemoveEntry(
				&whitelist.RemoveEntryParams{
					Signature: &tokenDetails.Signature,
				},
			)
			if errRemoveResult != nil {
				return nil, errRemoveResult
			}
			if removeResult == false {
				return nil, errRemoveResult
			}
			return entry, nil
		}
	}

	return nil, nil
}

func Update(p *UpdateParams) (*Session, error) {
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.SessionToken,
	)
	if errTokenDetails != nil {
		return nil, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Signature: &tokenDetails.Signature,
		},
	)
	if entry == nil {
		return nil, nil
	}
	if errEntry != nil {
		return nil, errEntry
	}

	resultJwt := jwtx.ValidateJWT(&jwtx.TokenPayload{
		Token:        tokenDetails,
		RandomSecret: &entry.SessionKey,
	})

	if resultJwt {
		sessionDetails, errSessionDetails := jwtx.RetrieveTokenDetailsFromString(
			p.SessionToken,
		)
		if errSessionDetails != nil {
			return nil, errSessionDetails
		}

		removeResult, errRemoveResult := whitelist.RemoveEntry(
			&whitelist.RemoveEntryParams{
				Signature: sessionDetails.Signature,
			},
		)
		if errRemoveResult != nil {
			return nil, errRemoveResult
		}
		if removeResult == false {
			return nil, errRemoveResult
		}

		issuedAt := jwtx.GetNowAsMS()
		expiresAt := issuedAt + getLifetimeByAudience(sessionDetails.Payload.Aud)

		return Create(&CreateParams{
			Iss: sessionDetails.Payload.Iss,
			Sub: sessionDetails.Payload.Sub,
			Aud: sessionDetails.Payload.Aud,
			Iat: issuedAt,
			Exp: expiresAt,
		})
	}

	return nil, nil
}

func Remove(p *RemoveParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
