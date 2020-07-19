package queries

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/hooks/verify"
	"webapi/sessions/sessionsx"
)

func isRequestBodyValid(
	w http.ResponseWriter,
	requestBody *requests.Body,
) bool {
	if requestBody != nil && requestBody.Params != nil {
		return true
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return false
}

func ValidateGuestSession(
	w http.ResponseWriter,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.ValidateGuest
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {		
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	sessionIsValid, errSessionIsValid := verify.ValidateGuestSession(
		params.Environment,
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: params.Token,
		},
	)
	if errSessionIsValid != nil {
		errors.DefaultResponse(w, errSessionIsValid)
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

func ValidateSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Validate
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {	
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	infraSessionIsValid, errInfraSessionIsValid := verify.ValidateInfraSession(
		params.Environment,
		sessionCookie,
	)
	if errInfraSessionIsValid != nil {
		errors.DefaultResponse(w, errInfraSessionIsValid)
		return
	}

	sessionIsValid, errSessionIsValid := sessionsx.Read(&params)
	if errSessionIsValid != nil {
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	
	if infraSessionIsValid && sessionIsValid {
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