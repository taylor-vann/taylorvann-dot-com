// brian taylor vann
// briantaylorvann dot com

// sessionx
//
// create, read, delete sessions
// read will automatically return expired sessions as invalid

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
	Aud 		 string	`json:"aud"`
	Iss 		 string	`json:"iss"`
	Lifetime int64	`json:"lifetime"`
	Sub 		 string	`json:"sub"`
}

type SessionClaims = jwtx.Claims

type UserParams struct {
	Environment string `json:"environment`
	UserID 			int64	 `json:"user_id"`
}

type CreateParams struct {
	Environment string 					`json:"environment`
	Claims 			*SessionClaims	`json:"claims"`
}

type ReadParams struct {
	Environment string `json:"environment`
	Token 			string `json:"token"`
}

type ValidateParams = ReadParams
type UpdateParams = ReadParams
type DeleteParams = whitelist.RemoveEntryParams

func CreateSessionClaims(p *CreateClaimsParams) *SessionClaims {
	issuedAt := jwtx.GetNowAsMS()
	expiresAt := issuedAt + p.Lifetime

	claims := SessionClaims{
		Iss: p.Iss,
		Sub: p.Sub,
		Aud: p.Aud,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &claims
}

func CreateGuestSessionClaims() *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: constants.Client,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeDaysAsMS,
		Sub: constants.Guest,
	})
}

func CreateInfraSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeSixtyFiveDaysAsMS,
		Sub: constants.Infra,
	})
}

func CreateClientSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeSixtyFiveDaysAsMS,
		Sub: constants.Client,
	})
}

func CreateCreateAccountSessionClaims(userID int64) (*SessionClaims) {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub: constants.CreateAccount,
	})
}

func CreateUpdateEmailSessionClaims(userID int64) (*SessionClaims) {	
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub: constants.UpdateEmail,
	})
}

func CreateUpdatePasswordSessionClaims(userID int64) (*SessionClaims) {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub: constants.UpdatePassword,
	})
}

func CreateDeleteAccountSessionClaims(userID int64) (*SessionClaims) {	
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud: userIDAsStr,
		Iss: constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub: constants.DeleteAccount,
	})
}

func Create(p *CreateParams) (*Session, error) {
	if p == nil {
		return nil, errors.New("nil params provided")
	}
	if p.Claims == nil {
		return nil, errors.New("nil claims provided")
	}

	token, errToken := jwtx.CreateJWT(p.Claims)
	if errToken != nil {
		return nil, errToken
	}

	lifetime := p.Claims.Exp - p.Claims.Iat
	if lifetime < 0 {
		return nil, errors.New("negative expiration provided")
	}
	_, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			Environment: p.Environment,
			SessionKey:	 token.RandomSecret,
			Signature:   token.Token.Signature,
			CreatedAt:   p.Claims.Iat,
			Lifetime:    lifetime,
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

	isGenericallyValid := jwtx.ValidateGenericToken(
		&jwtx.ValidateGenericTokenParams{
			Token: p.Token,
			Issuer: constants.BrianTaylorVannDotCom,
		},
	)
	if !isGenericallyValid {
		return false, errors.New("token is generically invalid")
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenFromString(
		p.Token,
	)
	if tokenDetails == nil || errTokenDetails != nil {
		return false, errTokenDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Environment: p.Environment,
			Signature:	 tokenDetails.Signature,
		},
	)
	if entry == nil || errEntry != nil {
		return false, errEntry
	}

	isSignedJWT := jwtx.ValidateJWT(&jwtx.TokenPayload{
		Token:        *tokenDetails,
		RandomSecret: entry.SessionKey,
	})

	return isSignedJWT, nil
}

func Delete(p *DeleteParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
