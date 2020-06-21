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

const SessionCookieHeader = "briantaylorvann_session"

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.CustomResponse(w, errors.BadRequestFail)
	return true
}

func dropRequestNotValidInfraSession(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return true
	}
	isValid, errValidate := verify.ValidateInfraSession(
		environment,
		sessionCookie.Value,
	)
	if isValid {
		return false
	}

	errors.DefaultResponse(w, errValidate)
	return true
}

func ValidateGuestSession(
	w http.ResponseWriter,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
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
		params.Token,
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
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	var params requests.Validate
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {	
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	if dropRequestNotValidInfraSession(w, params.Environment, sessionCookie) {
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