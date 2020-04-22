// brian taylor vann
// taylorvann dot com

package sessionsx

import (
	"errors"
	"webapi/interfaces/jwtx"
	"webapi/sessions/constants"
	"webapi/sessions/whitelist"
)

type MilliSeconds = int64

type Session struct {
	SessionToken string `json:"session_token"`
}

type CreateClaimsParams struct {
	Iss string
	Sub string
	Aud string
}

type SessionClaims = jwtx.Claims

type CreateUserClaimsParams struct {
	UserID int64
}

type CreateUserAccountClaimsParams struct {
	Email string
}

type CreateParams struct {
	Environment string
	Claims 			SessionClaims
}

type ReadParams struct {
	Environment string
	SessionToken string
}

type UpdateParams struct {
	Environment  string
	SessionToken string
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

func CreateSessionClaims(p *CreateClaimsParams) *SessionClaims {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(constants.Guest)

	claims := SessionClaims{
		Iss: p.Iss,
		Sub: p.Sub,
		Aud: p.Aud,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &claims
}

func CreateDocumentSessionClaims() *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: constants.Guest,
		Aud: constants.Document,
	})
}

func CreateGuestSessionClaims() *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: constants.Guest,
		Aud: constants.Public,
	})
}

func CreateUpdatePasswordSessionClaims(p *CreateUserAccountClaimsParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.UpdatePassword,
	})

	return claims, nil
}

func CreateUpdateEmailSessionClaims(p *CreateUserAccountClaimsParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.UpdateEmail,
	})

	return claims, nil
}

func CreateDeleteAccountSessionClaims(p *CreateUserAccountClaimsParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.DeleteAccount,
	})

	return claims, nil
}

func CreateAccountCreationSessionClaims(p *CreateUserAccountClaimsParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.CreateAccount,
	})

	return claims, nil
}

func CreateUserSessionClaims(p *CreateUserClaimsParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.TaylorVannDotCom,
		Sub: string(p.UserID),
		Aud: constants.Public,
	})

	return claims, nil
}

func Create(p *CreateParams) (*Session, error) {
	if p == nil {
		return nil, errors.New("nil CreateParams provided")
	}

	token, errToken := jwtx.CreateJWT(&p.Claims)
	if errToken != nil {
		return nil, errToken
	}

	_, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			Environment: p.Environment,
			SessionKey:	 token.RandomSecret,
			Signature:   token.Token.Signature,
			CreatedAt:   p.Claims.Iat,
			Lifetime:    getLifetimeByAudience(p.Claims.Aud),
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	sessionTokenAsStr, errSessionTokenAsStr := jwtx.ConvertTokenToString(
		&token.Token,
	)
	if errSessionTokenAsStr != nil {
		return nil, errSessionTokenAsStr
	}

	session := Session{
		SessionToken: sessionTokenAsStr,
	}

	return &session, nil
}

func Read(p *ReadParams) (bool, error) {
	if p == nil {
		return false, errors.New("nil params")
	}

	if p.SessionToken == "" {
		return false, errors.New("nil session token provided")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.SessionToken,
	)
	if errTokenDetails != nil {
		return false, errTokenDetails
	}
	if tokenDetails == nil {
		return false, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Environment: p.Environment,
			Signature:	 tokenDetails.Signature,
		},
	)
	if errEntry != nil {
		return false, errEntry
	}

	if entry != nil {
		result := jwtx.ValidateJWT(&jwtx.TokenPayload{
			Token:        *tokenDetails,
			RandomSecret: entry.SessionKey,
		})
		return result, nil
	}

	return false, nil
}

func ValidateAndRemove(p *ValidateAndRemoveParams) (*whitelist.Entry, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	if p.SessionToken == "" {
		return nil, errors.New("nil sesion provided")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.SessionToken,
	)
	if errTokenDetails != nil {
		return nil, errTokenDetails
	}
	if tokenDetails == nil {
		return nil, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Environment: p.Environment,
			Signature: tokenDetails.Signature,
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	if entry != nil {
		resultJwt := jwtx.ValidateJWT(&jwtx.TokenPayload{
			Token:        *tokenDetails,
			RandomSecret: entry.SessionKey,
		})
		if resultJwt {
			removeResult, errRemoveResult := whitelist.RemoveEntry(
				&whitelist.RemoveEntryParams{
					Environment: p.Environment,
					Signature: tokenDetails.Signature,
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
	if tokenDetails == nil {
		return nil, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Environment: p.Environment,
			Signature: 	 tokenDetails.Signature,
		},
	)
	if entry == nil {
		return nil, nil
	}
	if errEntry != nil {
		return nil, errEntry
	}

	resultJwt := jwtx.ValidateJWT(&jwtx.TokenPayload{
		Token:        *tokenDetails,
		RandomSecret: entry.SessionKey,
	})

	if resultJwt {
		sessionDetails, errSessionDetails := jwtx.RetrieveTokenDetailsFromString(
			p.SessionToken,
		)
		if errSessionDetails != nil {
			return nil, errSessionDetails
		}

		issuedAt := jwtx.GetNowAsMS()
		expiresAt := issuedAt + getLifetimeByAudience(sessionDetails.Payload.Aud)

		return Create(&CreateParams{
			Environment: p.Environment,
			Claims: SessionClaims{
				Iss: sessionDetails.Payload.Iss,
				Sub: sessionDetails.Payload.Sub,
				Aud: sessionDetails.Payload.Aud,
				Iat: issuedAt,
				Exp: expiresAt,
			},
		})
	}

	return nil, nil
}

func Remove(p *RemoveParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
