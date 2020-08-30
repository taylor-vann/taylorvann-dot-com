package verify

import (
	"errors"
	"net/http"

	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const (
	BrianTaylorVannDotCom = "briantaylorvann.com"

	Guest  = "guest"
	Client = "client"
	Infra  = "infra"
)

var (
	errSessionCookieIsNil    = errors.New("session cookie is nil")
	errGuestSessionIsInvalid = errors.New("guest session was invalid")
	errInfraSessionIsInvalid = errors.New("infra session was invalid")
)

func CheckGuestSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  BrianTaylorVannDotCom,
		Subject: Guest,
	})
}

func CheckClientSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  BrianTaylorVannDotCom,
		Subject: Client,
	})
}

func CheckInfraSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  BrianTaylorVannDotCom,
		Subject: Infra,
	})
}

func ValidateGuestSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errSessionCookieIsNil
	}

	isValid := CheckGuestSession(sessionCookie.Value)
	if !isValid {
		return false, errGuestSessionIsInvalid
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token:       sessionCookie.Value,
	})
}

func ValidateInfraSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errSessionCookieIsNil
	}

	isValid := CheckInfraSession(sessionCookie.Value)
	if !isValid {
		return false, errInfraSessionIsInvalid
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token:       sessionCookie.Value,
	})
}

func ValidateClientSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errSessionCookieIsNil
	}

	isValid := CheckClientSession(sessionCookie.Value)
	if !isValid {
		return false, errInfraSessionIsInvalid
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token:       sessionCookie.Value,
	})
}
