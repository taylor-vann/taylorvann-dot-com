package verifysession

import (
	"errors"
	"net/http"

	"log"

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
	log.Println("check guest session")
	result := jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
		Subject: "guest",
	})
	log.Println(result)
	details, _ := jwtx.RetrieveTokenDetailsFromString(sessionToken)
	nowAsMS := jwtx.GetNowAsMS()
	log.Println(details.Payload.Iss)
	log.Println(details.Payload.Sub)
	log.Println(details.Payload.Aud)
	log.Println(details.Payload.Iat)
	log.Println(details.Payload.Exp)
	log.Println(details.Payload.Exp - details.Payload.Iat)
	log.Println(details.Payload.Sub == "guest")
	log.Println(details.Payload.Iss == "briantaylorvann.com")
	log.Println(details.Payload.Iat < nowAsMS)
	log.Println(details.Payload.Iat < details.Payload.Exp)
	log.Println(nowAsMS < details.Payload.Exp)
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
	log.Println("verifysession, validateGuestSession")
	log.Println(environment)
	log.Println(sessionToken)
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