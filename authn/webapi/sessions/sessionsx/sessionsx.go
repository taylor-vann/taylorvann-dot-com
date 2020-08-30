// brian taylor vann
// briantaylorvann dot com

package sessionsx

import (
	"errors"
	"strconv"

	"webapi/sessions/sessionsx/constants"
	"webapi/sessions/whitelist"

	"webapi/jwtx"
)

type MilliSeconds = int64

type Session struct {
	Token string `json:"token"`
}

type CreateClaimsParams struct {
	Aud      string `json:"aud"`
	Iss      string `json:"iss"`
	Lifetime int64  `json:"lifetime"`
	Sub      string `json:"sub"`
}

type SessionClaims = jwtx.Claims

type CreateParams struct {
	Environment string         `json:"environment`
	Claims      *SessionClaims `json:"claims"`
}

type ReadParams struct {
	Environment string `json:"environment`
	Token       string `json:"token"`
}

type ValidateParams = ReadParams
type UpdateParams = ReadParams
type DeleteParams = whitelist.RemoveEntryParams

var (
	errNilParamsProvided          = errors.New("nil params provided")
	errNilClaimsProvided          = errors.New("nil claims provided")
	errNilSessionTokenProvided    = errors.New("nil session token provided")
	errTokenIsGenericallyInvalid  = errors.New("token is generically invalid")
	errNegativeExpirationProvided = errors.New("negative expiration provided")
)

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
		Aud:      constants.Client,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeDaysAsMS,
		Sub:      constants.Guest,
	})
}

func CreateInfraSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      userIDAsStr,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeSixtyFiveDaysAsMS,
		Sub:      constants.Infra,
	})
}

func CreateClientSessionClaims(userID int64) *SessionClaims {
	userIDAsStr := strconv.FormatInt(userID, 10)
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      userIDAsStr,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeSixtyFiveDaysAsMS,
		Sub:      constants.Client,
	})
}

func CreateCreateAccountSessionClaims(email string) *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      email,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub:      constants.CreateAccount,
	})
}

func CreateUpdateEmailSessionClaims(email string) *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      email,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeHoursAsMS,
		Sub:      constants.UpdateEmail,
	})
}

func CreateUpdatePasswordSessionClaims(email string) *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      email,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.ThreeHoursAsMS,
		Sub:      constants.UpdatePassword,
	})
}

func CreateDeleteAccountSessionClaims(email string) *SessionClaims {
	return CreateSessionClaims(&CreateClaimsParams{
		Aud:      email,
		Iss:      constants.BrianTaylorVannDotCom,
		Lifetime: constants.OneDayAsMS,
		Sub:      constants.DeleteAccount,
	})
}

func Create(p *CreateParams) (*Session, error) {
	if p == nil {
		return nil, errNilParamsProvided
	}
	if p.Claims == nil {
		return nil, errNilClaimsProvided
	}

	token, errToken := jwtx.CreateJWT(p.Claims)
	if errToken != nil {
		return nil, errToken
	}

	lifetime := p.Claims.Exp - p.Claims.Iat
	if lifetime < 0 {
		return nil, errNegativeExpirationProvided
	}
	_, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			Environment: p.Environment,
			SessionKey:  token.RandomSecret,
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
		return false, errNilParamsProvided
	}
	if p.Token == "" {
		return false, errNilSessionTokenProvided
	}

	isGenericallyValid := jwtx.ValidateGenericToken(
		&jwtx.ValidateGenericTokenParams{
			Token:  p.Token,
			Issuer: constants.BrianTaylorVannDotCom,
		},
	)
	if !isGenericallyValid {
		return false, errTokenIsGenericallyInvalid
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
			Signature:   tokenDetails.Signature,
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
