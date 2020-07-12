package verify

import (
	"errors"

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

func ValidateGuestSession(environment string, sessionToken string) (bool, error) {
	isValid := CheckGuestSession(sessionToken)
	if !isValid {
		return false, errors.New("guest session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionToken,
	})
}

func ValidateInfraSession(environment string, sessionToken string) (bool, error) {
	isValid := CheckInfraSession(sessionToken)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionToken,
	})
}

func ValidateClientSession(environment string, sessionToken string) (bool, error) {
	isValid := CheckClientSession(sessionToken)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionToken,
	})
}