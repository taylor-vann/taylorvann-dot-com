package queries

import (
	"encoding/json"
	"net/http"


	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/toolbox-go/jwtx"
)

const SessionCookieHeader = "briantaylorvann_session"

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
}

func checkGuestSession(sessionToken string) (bool, error) {
	isValid := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
	})
	if !isValid {
		return false, nil
	}

	details, errDetails := jwtx.RetrieveTokenDetailsFromString(sessionToken)
	if errDetails != nil {
		return false, errDetails
	}

	if details.Payload.Sub == "guest" && details.Payload.Aud == "client" {
		return true, nil
	}

	return false, nil
}

func checkInfraSession(sessionToken string) (bool, error) {
	isValid := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
	})
	if !isValid {
		return false, nil
	}

	details, errDetails := jwtx.RetrieveTokenDetailsFromString(sessionToken)
	if errDetails != nil {
		return false, errDetails
	}

	if details.Payload.Sub == "infra" {
		return true, nil
	}

	return false, nil
}

func validateSessionInWhitelist(environment string, sessionToken string) (bool, error) {
	return sessionsx.Read(&sessionsx.ValidateParams{
		Environment: environment,
		Token: sessionToken,
	})
}

func validateGuestSessionCache(environment string, sessionToken string) (bool, error) {
	guestSessionExists, errGuestSessionExists := validateSessionInWhitelist(
		environment,
		sessionToken,
	)
	if errGuestSessionExists != nil {
		return false, errGuestSessionExists
	}
	if !guestSessionExists {
		return false, nil
	}

	return checkGuestSession(sessionToken)
}

func validateInfraSessionCache(environment string, sessionToken string) (bool, error) {
	infraSessionExists, errInfraSessionExists := validateSessionInWhitelist(
		environment,
		sessionToken,
	)
	if errInfraSessionExists != nil {
		return false, errInfraSessionExists
	}
	if !infraSessionExists {
		return false, nil
	}

	return checkInfraSession(sessionToken)
}

// requires no cookies or authorization
func ValidateGuestSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.ValidateGuest
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {		
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateGuestSessionCache(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil {
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidGuestCredentials)
		return
	}
	
	if sessionIsValid {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&responses.Body{
			Session: &responses.Session{
				Token: sessionCookie.Value,
			},
		})
		return
	}

	errors.CustomResponse(w, errors.InvalidSessionCredentials)
}

func ValidateSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	// check session
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}
	
	// unmarshal params
	var params requests.Validate
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {	
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	infraSessionIsValid, errInfraSessionIsValid := validateInfraSessionCache(
		params.Environment,
		sessionCookie.Value,
	)
	if errInfraSessionIsValid != nil {
		errors.DefaultResponse(w, errInfraSessionIsValid)
		return
	}
	if !infraSessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&params)
	if errReadSession != nil {
		errors.DefaultResponse(w, errReadSession)
		return
	}
	
	if sessionIsValid {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&responses.Body{
			Session: &responses.Session{
				Token: params.Token,
			},
		})
		return
	}

	errors.CustomResponse(w, errors.InvalidSessionCredentials)
}