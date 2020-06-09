// brian taylor vann
// taylorvann dot com

package sessionsx

import (
	"errors"
	"strconv"
	
	"webapi/sessions/sessionsx/constants"
	"webapi/sessions/whitelist"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

type MilliSeconds = int64

type Session struct {
	Token string `json:"token"`
}

type CreateClaimsParams struct {
	Iss string	`json:"iss"`
	Sub string	`json:"sub"`
	Aud string	`json:"aud"`
}

type SessionClaims = jwtx.Claims

type UserParams struct {
	Environment string `json:"environment`
	UserID 			int64	 `json:"user_id"`
}

type AccountParams struct {
	Environment string 	`json:"environment`
	Email 			string	`json:"email"`
}

type CreateParams struct {
	Environment string 				`json:"environment`
	Claims 			SessionClaims	`json:"claims"`
}

type ReadParams struct {
	Environment string `json:"environment`
	Token 			string `json:"token"`
}

type ValidateParams = ReadParams
type UpdateParams = ReadParams

type DeleteParams = whitelist.RemoveEntryParams

func getLifetimeByAudience(audience string) int64 {
	switch audience {
	case constants.Guest:
		return constants.OneDayAsMS
	case constants.Client:
		return constants.ThreeDaysAsMS
	case constants.Infra:
		return constants.ThreeSixtyFiveDaysAsMS
	default:
		return constants.OneDayAsMS
	}
}

func CreateSessionClaims(p *CreateClaimsParams) *SessionClaims {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + getLifetimeByAudience(p.Sub)

	claims := SessionClaims{
		Iss: p.Iss,
		Sub: p.Sub,
		Aud: p.Aud,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &claims
}

func CreateInfraSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: constants.Infra,
		Aud: userIDAsStr,
	})
}

func CreateClientSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: constants.Client,
		Aud: userIDAsStr,
	})
}

func CreateGuestSessionClaims() *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: constants.Guest,
		Aud: constants.Client,
	})
}

func CreateUpdatePasswordSessionClaims(p *AccountParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.UpdatePassword,
	})

	return claims, nil
}

func CreateUpdateEmailSessionClaims(p *AccountParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.UpdateEmail,
	})

	return claims, nil
}

func CreateDeleteAccountSessionClaims(p *AccountParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}
	
	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.DeleteAccount,
	})

	return claims, nil
}

func CreateAccountCreationSessionClaims(p *AccountParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: p.Email,
		Aud: constants.CreateAccount,
	})

	return claims, nil
}

func CreateUserSessionClaims(p *UserParams) (*SessionClaims, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	claims := CreateSessionClaims(&CreateClaimsParams{
		Iss: constants.BrianTaylorVannDotCom,
		Sub: string(p.UserID),
		Aud: constants.Client,
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
		Token: sessionTokenAsStr,
	}

	return &session, nil
}

func Read(p *ReadParams) (bool, error) {
	if p == nil {
		return false, errors.New("nil params")
	}

	if p.Token == "" {
		return false, errors.New("nil session token provided")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.Token,
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

func Update(p *UpdateParams) (*Session, error) {
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.Token,
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
			p.Token,
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

func Delete(p *DeleteParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
