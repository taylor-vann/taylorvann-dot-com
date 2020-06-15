package verifysession

import (
	"errors"
	"net/http"

	hookErrors "webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const SessionCookieHeader = "briantaylorvann_session"

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	hookErrors.BadRequest(w, &responses.Errors{
		RequestBody: &hookErrors.BadRequestFail,
	})
	return true
}

func CheckGuestSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
		Subject: "guest",
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