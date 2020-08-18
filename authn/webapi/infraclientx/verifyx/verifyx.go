package verifyx

import (
	"net/http"

	"webapi/infraclientx/fetchx/requests"
	"webapi/infraclientx/fetchx"
	"webapi/infraclientx/verifyx/errors"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const issuer = "briantaylorvann.com"

func CheckGuestSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: issuer,
		Subject: "guest",
	})
}

func CheckClientSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: issuer,
		Subject: "client",
	})
}

func CheckInfraSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token: sessionToken,
		Issuer: issuer,
		Subject: "infra",
	})
}

func IsGuestSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	if sessionCookie == nil {
		return false
	}
	if !CheckGuestSession(sessionCookie.Value) {
		return false
	}
	validToken, errValidate := fetchx.ValidateGuestSession(
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

func IsInfraSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionCookie *http.Cookie,
) bool {
	if sessionCookie == nil {
		return false
	}
	if !CheckInfraSession(sessionCookie.Value) {
		return false
	}
	validToken, errValidToken := fetchx.ValidateSession(
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

func IsSessionValid(
	w http.ResponseWriter,
	environment string,
	infraSessionCookie *http.Cookie,
	sessionToken string,
) bool {
	if infraSessionCookie == nil {
		return false
	}

	validToken, errValidToken := fetchx.ValidateSession(
		&requests.ValidateSession{
			Environment: environment,
			Token: sessionToken,
		},
		infraSessionCookie,
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

// has role from session
func HasRoleFromSession(
	w http.ResponseWriter,
	environment string,
	infraSessionCookie *http.Cookie,
	sessionToken string,
	organization string,
) bool {
	if infraSessionCookie == nil {
		return false
	}

	validRole, errValidRole := fetchx.ValidateRoleFromSession(
		&requests.ValidateRoleFromSession{
			Environment: environment,
			Token: sessionToken,
			Organization: organization,
		},
		infraSessionCookie,
	)
	if errValidRole != nil {
		errors.DefaultResponse(w, errValidRole)
		return false
	}
	if validRole != nil {
		return true
	}
	
	errors.CustomResponse(w, errors.InvalidInfraSession)
	return false
}

// validate user
func ValidateUser(
	w http.ResponseWriter,
	environment string,
	infraSessionCookie *http.Cookie,
	email string,
	password string,
) bool {
	if infraSessionCookie == nil {
		return false
	}
	validRole, errValidRole := fetchx.ValidateUser(
		&requests.ValidateUser{
			Environment: environment,
			Email: email,
			Password: password,
		},
		infraSessionCookie,
	)
	if validRole != nil {
		return true
	}
	if errValidRole != nil {
		errors.DefaultResponse(w, errValidRole)
		return false
	}
	
	errors.CustomResponse(w, errors.InvalidInfraSession)
	return false
}