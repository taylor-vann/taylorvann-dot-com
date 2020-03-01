// brian taylor vann
// taylorvann dot com
//
// Package sessions - meta-interface for whitelist, csrfs, and session tokens

package sessions

import (
	"errors"
	"time"
	"webapi/interfaces/jwtx"
	"webapi/sessions/constants"
	"webapi/store"
	"webapi/whitelist"
)

// MilliSeconds -
type MilliSeconds = int64

// Session -
type Session struct {
	SessionToken string	`json:"session_token"`
	CsrfToken    []byte	`json:"csrf_token"`
}

// CreatePublicJWTParams -
type CreatePublicJWTParams struct {
	Email    string
	Password string
}

// CreateParams -
type CreateParams struct {
	Issuer   string
	Subject  string
	Audience string
}

// ReadParams -
type ReadParams struct {
	SessionToken	*string
}

// CheckParams -
type CheckParams struct {
	SessionToken	*string
	CsrfToken			*[]byte
}

// UpdateParams -
type UpdateParams = CheckParams

// RemoveParams -
type RemoveParams = whitelist.RemoveEntryParams

// getNowAsMS -
func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func composeJWTClaims(p *CreateParams, issuedAt int64, expiresAt int64) *jwtx.Claims {
	jwtClaims := jwtx.Claims{
		Iss: p.Issuer,
		Sub: p.Subject,
		Aud: p.Audience,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &jwtClaims
}

func validateCsrfTokens(external *[]byte, whitelist *[]byte) bool {
	if len(*external) == len(*whitelist) {
		for index, bit := range *whitelist {
			if (*external)[index] != bit {
				return false
			}
		}

		return true
	}

	return false
}

// ComposeCreateGuestSessionParams -
func ComposeCreateGuestSessionParams() *CreateParams {
	params := CreateParams{
		Issuer:   constants.TaylorVannDotCom,
		Subject:  constants.Guest,
		Audience: constants.Public,
	}

	return &params
}

// ComposeCreateResetPasswordSessionParams -
func ComposeCreateResetPasswordSessionParams() *CreateParams {
	params := CreateParams{
		Issuer:   constants.TaylorVannDotCom,
		Subject:  constants.ResetPassword,
		Audience: constants.Public,
	}

	return &params
}

// ComposeCreatePublicSessionParams - validate user through store
func ComposeCreatePublicSessionParams(p *CreatePublicJWTParams) (*CreateParams, error) {
	userRow, errValidUser := store.ValidateUser(
		&store.ValidateUserParams{
			Email:    p.Email,
			Password: p.Password,
		},
	)
	if userRow == nil {
		return nil, errors.New("bad credentials provided")
	}
	if errValidUser != nil {
		return nil, errValidUser
	}

	params := CreateParams{
		Issuer:   constants.TaylorVannDotCom,
		Subject:  string(userRow.ID),
		Audience: constants.Public,
	}

	return &params, nil
}

// Create -
func Create(p *CreateParams) (*Session, error) {
	if p == nil {
		return nil, errors.New("nil CreateParams provided")
	}
	
	// create guest jwt
	issuedAt := getNowAsMS()
	lifetime := constants.ThreeDaysAsMS
	expiresAt := issuedAt + lifetime

	token, errToken := jwtx.CreateJWT(
		composeJWTClaims(p, issuedAt, expiresAt),
	)
	if errToken != nil {
		return nil, errToken
	}

	// add entry
	entry, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			CreatedAt:  issuedAt,
			Lifetime:   lifetime,
			SessionKey: token.RandomSecret,
			Signature:  &token.Token.Signature,
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	// compose session
	sessionTokenAsStr, errSessionTokenAsStr := jwtx.ConvertTokenToString(
		token.Token,
	)
	if errSessionTokenAsStr != nil {
		return nil, errSessionTokenAsStr
	}

	session := Session{
		SessionToken: *sessionTokenAsStr,
		CsrfToken:    entry.CsrfToken,
	}

	return &session, nil
}

// Read -
func Read(p *ReadParams) (bool, error) {
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

	result := jwtx.ValidateJWT(&jwtx.TokenPayload{
		Token: tokenDetails,
		RandomSecret: &entry.SessionKey,
	})

	return result, nil
}

// Check -
func Check(p *CheckParams) (*whitelist.Entry, error) {
	if p.SessionToken == nil {
		return nil, nil
	}
	if p.CsrfToken == nil {
		return nil, nil
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
			Token: tokenDetails,
			RandomSecret: &entry.SessionKey,
		})
		resultCsrf := validateCsrfTokens(p.CsrfToken, &entry.CsrfToken)
		if resultJwt && resultCsrf {
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

// Update -
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
		Token: tokenDetails,
		RandomSecret: &entry.SessionKey,
	})
	resultCsrf := validateCsrfTokens(p.CsrfToken, &entry.CsrfToken)

	if resultJwt && resultCsrf {
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

		return Create(&CreateParams{
			Issuer:   sessionDetails.Payload.Iss,
			Subject:  sessionDetails.Payload.Sub,
			Audience: sessionDetails.Payload.Aud,
		})
	}

	return nil, nil
}

// Remove -
func Remove(p *RemoveParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
