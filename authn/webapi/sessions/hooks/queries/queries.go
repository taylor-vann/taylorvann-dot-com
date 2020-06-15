package queries

import (
	"encoding/json"
	"net/http"

	"log"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
	"webapi/sessions/hooks/verifysession"
)

const SessionCookieHeader = "briantaylorvann_session"

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.CustomResponse(w, errors.BadRequestFail)
	return true
}

// dropRequestNotValidInfraSession
func dropRequestNotValidInfraSession(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return true
	}
	isValid, errValidate := verifysession.ValidateInfraSession(
		environment,
		sessionCookie.Value,
	)
	if isValid {
		return false
	}

	errors.DefaultResponse(w, errValidate)
	return true
}

// requires no cookies or authorization
func ValidateGuestSession(
	w http.ResponseWriter,
	requestBody *requests.Body,
) {
	log.Println("ValidateGuestSession")
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	log.Println("request body was valid")
	var params requests.ValidateGuest
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsUnmarshal := json.Unmarshal(bytes, &params)
	if errParamsUnmarshal != nil {		
		log.Println(errParamsUnmarshal)
		errors.DefaultResponse(w, errParamsUnmarshal)
		return
	}

	sessionIsValid, errSessionIsValid := verifysession.ValidateGuestSession(
		params.Environment,
		params.Token,
	)
	if errSessionIsValid != nil {
		log.Println("session was invalid")

		log.Println(errSessionIsValid)
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	
	if sessionIsValid {
		log.Println("valid session!")

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
	// check session
	if dropRequestNotValidBody(w, requestBody) {
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

	// check session
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