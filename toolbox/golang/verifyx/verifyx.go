package verifyx

import (
	"errors"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/requests"
	"github.com/taylor-vann/weblog/toolbox/golang/clientx"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const SessionCookieHeader = "briantaylorvann_session"

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

func RemotelyValidateSession(environment string, sessionToken string) (bool, error) {
	sessionStr, errSessionStr := clientx.ValidateSession(
		requests.ValidateSession{
			Environment: environment,
			Token: sessionToken,
		},
	)
	if errSessionStr != nil {
		return false, errSessionStr
	}
	if sessionStr != "" {
		return true, nil
	}

	return false, errors.New("fail automatically")
}

func ValidateGuestSession(environment string, sessionToken string) (bool, error) {
	isValid := CheckGuestSession(sessionToken)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return RemotelyValidateSession(environment, sessionToken)
}

func ValidateInfraSession(environment string, sessionToken string) (bool, error) {
	isValid := CheckInfraSession(sessionToken)
	if !isValid {
		return false, errors.New("infra session was invalid")
	}

	return RemotelyValidateSession(environment, sessionToken)
}