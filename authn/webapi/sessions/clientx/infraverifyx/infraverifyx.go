package infraverifyx

import (
	"net/http"

	"webapi/sessions/clientx/infrafetchx/requests"
	"webapi/sessions/clientx/infrafetchx"
	"webapi/sessions/clientx/infraverifyx/errors"
)

func IsGuestSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	validToken, errValidate := infrafetchx.ValidateGuestSession(
		&requests.ValidateSession{
			Environment: environment,
			Token: sessionCookie.Value,
		},
		sessionCookie,
	)
	if validToken != nil {
		return true
	}
	if errValidate != nil {
		errors.DefaultResponse(w, errValidate)
		return false
	}
	
	errors.CustomResponse(w, errors.InvalidGuestSession)
	return false
}

func IsSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	validToken, errValidToken := infrafetchx.ValidateSession(
		&requests.ValidateSession{
			Environment: environment,
			Token: sessionCookie.Value,
		},
		sessionCookie,
	)
	if validToken != nil {
		return true
	}
	if errValidToken != nil {
		errors.DefaultResponse(w, errValidToken)
		return false
	}
	
	errors.CustomResponse(w, errors.InvalidInfraSession)
	return false
}