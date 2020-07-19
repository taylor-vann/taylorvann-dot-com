package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/infraclientx/fetchx"
	fetchRequests "webapi/sessions/infraclientx/fetchx/requests"
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/hooks/verify"
	"webapi/sessions/sessionsx"
)

const (
	ContentType = "Content-Type"
	ApplicationJson = "application/json"
	SessionCookieHeader = "briantaylorvann.com_session"
	CookieDomain = ".briantaylorvann.com"
	ThreeDaysInSeconds = 60 * 60 * 24 * 3
	ThreeSixtyFiveDaysInSeconds = 60 * 60 * 24 * 365
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

func CreateGuestSession(
	w http.ResponseWriter,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
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
		Claims: sessionsx.CreateGuestSessionClaims(),
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	
	var params fetchRequests.ValidateGuestUser
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}
	
	// golang defaults strings to a "", exciting! (dangerous)
	if params.Email == "" {
		errors.CustomResponse(w, errors.InvalidDefaultUserProvided)
		return
	}

	resp, errResp := fetchx.ValidateInfraRole(
		&params,
		sessionCookie,
	)
	if errResp != nil {
		errors.DefaultResponse(w, errResp)
		return
	}
	
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: sessionsx.CreateInfraSessionClaims(resp.UserID),
	})
	if errSession == nil {
		http.SetCookie(w, createInfraSessionCookie(session.Token))
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateClientSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
		
	var params requests.User
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}
	
	// golang defaults ints to a 0 value, exciting! (dangerous)
	if params.UserID == 0 {
		errors.CustomResponse(w, errors.InvalidDefaultUserProvided)
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

	claims := sessionsx.CreateClientSessionClaims(params.UserID)
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: claims,
	})
	if infraSessionIsValid && errSession == nil {
		http.SetCookie(w, createClientSessionCookie(session.Token))
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateCreateAccountSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Ancillary
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	claims := sessionsx.CreateCreateAccountSessionClaims(params.Email)
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: claims,
	})
	if infraSessionIsValid && errSession == nil {
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateUpdateEmailSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Ancillary
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	claims := sessionsx.CreateUpdateEmailSessionClaims(params.Email)
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: claims,
	})
	if infraSessionIsValid && errSession == nil {
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateUpdatePasswordSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Ancillary
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	claims := sessionsx.CreateUpdatePasswordSessionClaims(params.Email)
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: claims,
	})
	if infraSessionIsValid && errSession == nil {
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func CreateDeleteAccountSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Ancillary
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	claims := sessionsx.CreateDeleteAccountSessionClaims(params.Email)
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: claims,
	})
	if infraSessionIsValid && errSession == nil {
		w.Header().Set(ContentType, ApplicationJson)
		json.NewEncoder(w).Encode(&responses.Body{
			Session: session,
		})
		return
	}

	errors.DefaultResponse(w, errSession)
}

func DeleteSession(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Delete
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	sessionWasDeleted, errSessionWasDeleted := sessionsx.Delete(&params)
	if errSessionWasDeleted != nil {
		errors.DefaultResponse(w, errSessionWasDeleted)
		return
	}
	if infraSessionIsValid && sessionWasDeleted {
		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusOK)
		return
	}

	errors.CustomResponse(w, errors.InvalidSessionProvided)
}
