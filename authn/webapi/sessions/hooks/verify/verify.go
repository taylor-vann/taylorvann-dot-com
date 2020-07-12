package verify

import (
	"errors"
	"net/http"

	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const SessionCookieHeader = "briantaylorvann.com_session"

func CheckGuestSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
		Subject: "guest",
	})
}

func CheckClientSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
		Subject: "client",
	})
}

func CheckInfraSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
		Subject: "infra",
	})
}

func ValidateGuestSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errors.New("session cookie is nil")
	}

	isValid := CheckGuestSession(sessionCookie.Value)
	if !isValid {
		return false, errors.New("guest session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionCookie.Value,
	})
}

func ValidateInfraSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errors.New("session cookie is nil")
	}

	isValid := CheckInfraSession(sessionCookie.Value)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionCookie.Value,
	})
}

func ValidateClientSession(environment string, sessionCookie *http.Cookie) (bool, error) {
	if sessionCookie == nil {
		return false, errors.New("session cookie is nil")
	}

	isValid := CheckClientSession(sessionCookie.Value)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionCookie.Value,
	})
}