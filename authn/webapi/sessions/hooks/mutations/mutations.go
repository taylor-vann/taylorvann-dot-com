package mutations

import (
	"encoding/json"
	"net/http"

	"github.com/weblog/toolbox/golang/clientx/fetch"
	fetchRequests "github.com/weblog/toolbox/golang/clientx/fetch/requests"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"

	"github.com/weblog/toolbox/golang/jwtx"
)

const (
	ContentType = "Content-Type"
	ApplicationJson = "application/json"
	SessionCookieHeader = "briantaylorvann.com_session"
	CookieDomain = ".briantaylorvann.com"
	ThreeDaysInSeconds = 60 * 60 * 24 * 3
	ThreeSixtyFiveDaysInSeconds = 60 * 60 * 24 * 365
)

func createGuestSessionCookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:			SessionCookieHeader,
		Value:		session,
		MaxAge:		ThreeDaysInSeconds,
		Domain:   CookieDomain,
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}

func createInfraSessionCookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:			SessionCookieHeader,
		Value:		session,
		MaxAge:		ThreeSixtyFiveDaysInSeconds,
		Domain:   CookieDomain,
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}

func createClientSessionCookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:			SessionCookieHeader,
		Value:		session,
		MaxAge:		ThreeDaysInSeconds,
		Domain:   CookieDomain,
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}



func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
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

func CreateGuestSession(w http.ResponseWriter, requestBody *requests.Body) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	
	var params requests.Guest
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionsx.CreateGuestSessionClaims(),
	})

	if errSession == nil {
		http.SetCookie(w, createGuestSessionCookie(session.Token))
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateInfraSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}
	
	var params requests.InfraUser
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	resp, errResp := fetch.ValidateInfraRole(
		fetchRequests.ValidateGuestUser(params),
		sessionCookie,
	)
	if errResp != nil {
		errors.DefaultResponse(w, errResp)
	}
	
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionsx.CreateInfraSessionClaims(resp.UserID),
	})

	if errSession == nil {
		http.SetCookie(w, createInfraSessionCookie(session.Token))
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
	}

	errors.DefaultResponse(w, errSession)
}

func CreateClientSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.InvalidInfraCredentials)
		return
	}
		
	var params requests.ClientUser
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	resp, errResp := fetch.ValidateGuestUser(
		fetchRequests.ValidateGuestUser(params),
		sessionCookie,
	)
	if errResp != nil {

		errors.DefaultResponse(w, errResp)
	}
	
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionsx.CreateClientSessionClaims(resp.ID),
	})

	if errSession == nil {
		http.SetCookie(w, createClientSessionCookie(session.Token))
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
	}

	errors.DefaultResponse(w, errSession)
}

// func updateGenericSession(p *requests.Update) (*sessionsx.Session, error) {
// 	if p == nil {
// 		return nil, err.New("request body is nil")
// 	}

// 	tokenResults := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
// 		Token:    p.Token,
// 		Issuer:		constants.TaylorVannDotCom,
// 	})
// 	if !tokenResults {
// 		return nil, err.New("unable to validate generic token")
// 	}

// 	session, errSession := sessionsx.Update(p)
// 	if errSession != nil {
// 		return nil, errSession
// 	}

// 	return session, nil
// }


// func CreateCreateAccountSession(w http.ResponseWriter, requestBody *requests.Body) {
// 	if dropRequestNotValidBody(w, requestBody) {
// 		return
// 	}

// 	bytes, _ := json.Marshal(requestBody.Params)
// 	var params requests.Account
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 			SessionToken: session.Token,
// 		})
// 		if errMarshal == nil {
// 			w.Header().Set(ContentType, ApplicationJson)
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
// 	var params requests.Account
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 				SessionToken: session.Token,
// 			},
// 		)
// 		if errMarshal == nil {
// 			w.Header().Set(ContentType, ApplicationJson)
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
// 	var params requests.Account
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 			SessionToken: session.Token,
// 		})
// 		if errMarshal == nil {
// 			w.Header().Set(ContentType, ApplicationJson)
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
// 	var params requests.User
// 	errParamsMarshal := json.Unmarshal(bytes, &params)
// 	if errParamsMarshal != nil {
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 				SessionToken: userSession.Token,
// 			},
// 		)
// 		if errMarshal == nil {
// 			w.Header().Set(ContentType, ApplicationJson)
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
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 				SessionToken: session.Token,
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

// 		w.Header().Set(ContentType, ApplicationJson)
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
// 		errors.DefaultResponse(w, errParamsMarshal)
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
// 		w.Header().Set(ContentType, ApplicationJson)
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	errors.CustomErrorResponse(w, errors.InvalidSessionProvided)
// }
