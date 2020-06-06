package queries

import (
	"encoding/json"
	"net/http"

	"log"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/tvgtb/jwtx"
)

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
}

func checkInfraSession(sessionCookie *http.Cookie) (bool, error) {
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(sessionCookie.Value)
	if errDetails != nil {
		return false, errDetails
	}
	if details == nil {
		return false, nil
	}
	if details.Payload.Sub == "infra" {
		return true, nil
	}

	return false, nil
}

func validateInfraSessionRemotely(w http.ResponseWriter, p *requests.Validate, sessionCookie *http.Cookie) (bool, error) {
	infraSessionExists, errInfraSessionExists := sessionsx.Read(&sessionsx.ValidateParams{
		Environment: p.Environment,
		Token: sessionCookie.Value,
	})
	if errInfraSessionExists != nil {
		log.Println("err checking infra session")
		return false, errInfraSessionExists
	}
	if !infraSessionExists {
		log.Println("infra session is not valid")
		return false, nil
	}

	return checkInfraSession(sessionCookie)
}

// requires no cookies or authorization
func ValidateGuestSession(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		sessionCookie.Value,
	)
	if errTokenDetails != nil {
		errors.DefaultResponse(w, errTokenDetails)
		return
	}
	if tokenDetails.Payload.Aud != "public" || tokenDetails.Payload.Sub != "guest" {
		errors.CustomResponse(w, errors.InvalidSessionCredentials)
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.ValidateGuest
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {		
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&requests.Read{
		Environment: params.Environment,
		Token: sessionCookie.Value,
	})
	if errReadSession != nil {
		errors.DefaultResponse(w, errReadSession)
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

func ValidateSession(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Validate
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {	
		log.Println("err params marshal")
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	infraSessionIsValid, errInfraSessionIsValid := validateInfraSessionRemotely(
		w,
		&params,
		sessionCookie,
	)
	if errInfraSessionIsValid != nil {
		log.Println("err with infra session")
		errors.DefaultResponse(w, errInfraSessionIsValid)
		return
	}
	if !infraSessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}

	// actual session
	sessionIsValid, errReadSession := sessionsx.Read(&params)
	if errReadSession != nil {
		log.Println("err reading session")
		errors.DefaultResponse(w, errReadSession)
		return
	}
	
	if sessionIsValid {
		log.Println("sessionIs valid!")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&responses.Body{
			Session: &responses.Session{
				Token: params.Token,
			},
		})
		return
	}

	log.Println("couldn't validate session")

	errors.CustomResponse(w, errors.InvalidSessionCredentials)
}