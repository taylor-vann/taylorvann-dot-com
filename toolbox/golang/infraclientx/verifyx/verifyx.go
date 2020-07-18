package verifyx

import (
	"net/http"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx/requests"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/verifyx/errors"

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
	sessionCookie *http.Cookie,
) bool {
	if sessionCookie == nil {
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