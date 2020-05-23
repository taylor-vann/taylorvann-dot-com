package mutations

import (
	"encoding/json"
	err "errors"
	// "io"
	"net/http"
	"log"
	
	"github.com/taylor-vann/tvgtb/jwtx"

	// "webapi/cookiesessionx"
	
	"webapi/sessions/hooks/constants"
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func dropRequestNotValidBody(
	w http.ResponseWriter,
	requestBody*requests.Body,
) bool {
	if requestBody != nil && requestBody.Params != nil {
		return true
	}
	errors.BadRequest(w, &responses.Errors{
		Body: &errors.BadRequestFail,
	})
	return false
}

func defaultErrorResponse(w http.ResponseWriter, errorResponse error) {
	errAsStr := errorResponse.Error()
	errors.BadRequest(w, &responses.Errors{
		Body: &errors.BadRequestFail,
		Default: &errAsStr,
	})
}

// side effects
func dropRequestUnableToMarshalBody(
	w http.ResponseWriter,
	requestBody *requests.Body,
	params interface{},
) bool {
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, params)
	if errParamsMarshal == nil {
		return true
	}
	defaultErrorResponse(w, errParamsMarshal)
	return false
}

func updateGenericSession(p *requests.Update) (*sessionsx.Session, error) {
	if p == nil {
		return nil, err.New("request body is nil")
	}

	tokenResults := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token:    p.SessionToken,
		Issuer:		constants.TaylorVannDotCom,
	})
	if !tokenResults {
		return nil, err.New("unable to validate generic token")
	}

	session, errSession := sessionsx.Update(p)
	if errSession != nil {
		return nil, errSession
	}

	return session, nil
}

// the only public mutation
func CreateGuestSession(w http.ResponseWriter, requestBody *requests.Body) {
	log.Println("made it to create guest session")	
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	var params requests.SessionParams
	if dropRequestUnableToMarshalBody(w, requestBody, params) {
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionsx.CreateGuestSessionClaims(),
	})

	if errSession == nil {
		log.Println("made it to the session")
		log.Println("session: " + session.SessionToken)

		// w.WriteHeader(http.StatusOK)
		// cookiesessionx.SetSessionCookie(w, session.SessionToken)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &responses.Errors{
		Session: &errors.CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}

// func CreateCreateAccountSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.AccountParams
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	sessionParams, errSessionParams := sessionsx.CreateAccountCreationSessionClaims(
// 		&params,
// 	)

// 	if errSessionParams != nil {
// 		errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
// 		return
// 	}

// 	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
// 		Environment: params.Environment,
// 		Claims: *sessionParams,
// 	})

// 	if errSession == nil {
// 		marshalledJSON, errMarshal := json.Marshal(&responses.Session{
// 			SessionToken: session.SessionToken,
// 		})
// 		if errMarshal == nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&marshalledJSON)
// 			return
// 		}

// 		errors.CustomErrorResponse(w, errors.UnableToMarshalSession)
// 		return
// 	}

// 	errorAsStr := errSession.Error()
// 	errors.BadRequest(w, &responses.Errors{
// 		Session: &errors.CreateGuestSessionErrorMessage,
// 		Default: &errorAsStr,
// 	})
// }

// func CreateUpdateEmailSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.AccountParams
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	userSessionToken, errUserSessionToken := sessionsx.CreateUpdatePasswordSessionClaims(
// 		&params,
// 	)
// 	if errUserSessionToken != nil {
// 		errorAsStr := errUserSessionToken.Error()
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.InvalidSessionCredentials,
// 			Default: &errorAsStr,
// 		})
// 		return
// 	}

// 	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
// 		Environment: params.Environment,
// 		Claims: *userSessionToken,
// 	})

// 	if errSession == nil {
// 		marshalledJSON, errMarshal := json.Marshal(
// 			&responses.Session{
// 				SessionToken: session.SessionToken,
// 			},
// 		)
// 		if errMarshal == nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&marshalledJSON)
// 			return
// 		}
		
// 		errors.CustomErrorResponse(w, errors.UnableToMarshalSession)
// 		return
// 	}

// 	errorAsStr := errSession.Error()
// 	errors.BadRequest(w, &responses.Errors{
// 		Session: &errors.CreateGuestSessionErrorMessage,
// 		Default: &errorAsStr,
// 	})
// }

// func CreateUpdatePasswordSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.AccountParams
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	userSessionToken, errUserSessionToken := sessionsx.CreateUpdatePasswordSessionClaims(
// 		&params,
// 	)
// 	if errUserSessionToken != nil {
// 		errorAsStr := errUserSessionToken.Error()
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.InvalidSessionCredentials,
// 			Default: &errorAsStr,
// 		})
// 		return
// 	}

// 	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
// 		Environment: params.Environment,
// 		Claims: *userSessionToken,
// 	})

// 	if errSession == nil {
// 		marshalledJSON, errMarshal := json.Marshal(&responses.Session{
// 			SessionToken: session.SessionToken,
// 		})
// 		if errMarshal == nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&marshalledJSON)
// 			return
// 		}

// 		errors.CustomErrorResponse(w, errors.UnableToMarshalSession)
// 		return
// 	}

// 	errorAsStr := errSession.Error()
// 	errors.BadRequest(w, &responses.Errors{
// 		Session: &errors.CreateGuestSessionErrorMessage,
// 		Default: &errorAsStr,
// 	})
// }


// func CreatePublicSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.UserParams
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	userSessionToken, errUserSessionToken := sessionsx.CreateUserSessionClaims(
// 		&params,
// 	)
// 	if errUserSessionToken != nil {
// 		errorAsStr := errUserSessionToken.Error()
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.InvalidSessionCredentials,
// 			Default: &errorAsStr,
// 		})
// 		return
// 	}

// 	userSession, errUserSession := sessionsx.Create(&sessionsx.CreateParams{
// 		Environment: params.Environment,
// 		Claims:	*userSessionToken,
// 	})

// 	if errUserSession == nil {
// 		marshalledJSON, errMarshal := json.Marshal(
// 			&responses.Session{
// 				SessionToken: userSession.SessionToken,
// 			},
// 		)
// 		if errMarshal == nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&marshalledJSON)
// 			return
// 		}

// 		errors.CustomErrorResponse(w, errors.UnableToMarshalSession)
// 		return
// 	}

// 	errorAsStr := errUserSession.Error()
// 	errors.BadRequest(w, &responses.Errors{
// 		Session: &errors.UnableToCreatePublicSession,
// 		Default: &errorAsStr,
// 	})
// }


// func UpdateSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if requestBody == nil || requestBody.Params == nil {
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.UnableToUpdateSession,
// 			Body: &errors.BadRequestFail,
// 		})
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.Update
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	session, errSession := updateGenericSession(&params)
// 	if errSession != nil {
// 		errAsStr := errSession.Error()
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.InvalidSessionProvided,
// 			Default: &errAsStr,
// 		})
// 		return
// 	}

// 	if errSession == nil {
// 		marshalledJSON, errMarshal := json.Marshal(
// 			&responses.Session{
// 				SessionToken: session.SessionToken,
// 			},
// 		)
// 		if errMarshal != nil {
// 			errAsStr := errMarshal.Error()
// 			errors.BadRequest(w, &responses.Errors{
// 				Session: &errors.UnableToMarshalSession,
// 				Default: &errAsStr,
// 			})
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(&marshalledJSON)
// 		return
// 	}

// 	errAsStr := errSession.Error()
// 	errors.BadRequest(w, &responses.Errors{
// 		Session: &errors.UnableToUpdateSession,
// 		Default: &errAsStr,
// 	})
// }

// func DeleteSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.Delete
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		defaultErrorResponse(w, errParamsMarshal)
// 		return
// 	}

// 	result, errResponseBody := sessionsx.Delete(&params)
// 	if errResponseBody != nil {
// 		errAsStr := errResponseBody.Error()
// 		errors.BadRequest(w, &responses.Errors{
// 			Session: &errors.InvalidSessionProvided,
// 			Default: &errAsStr,
// 		})
// 		return
// 	}

// 	if result == true {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	errors.CustomErrorResponse(w, errors.InvalidSessionProvided)
// }
