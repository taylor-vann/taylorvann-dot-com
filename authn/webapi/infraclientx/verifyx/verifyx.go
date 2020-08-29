package verifyx

import (
	"net/http"

	"webapi/infraclientx/fetchx"
	"webapi/infraclientx/fetchx/requests"
	"webapi/infraclientx/verifyx/errors"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

type ValidateUserParams struct {
	Environment        string       `json:"environment"`
	InfraSessionCookie *http.Cookie `json:"infra_session_cookie"`
	Email              string       `json:"email"`
	Password           string       `json:"password"`
}

type IsSessionValidParams struct {
	Environment        string       `json:"environment"`
	InfraSessionCookie *http.Cookie `json:"infra_session_cookie"`
	SessionCookie      *http.Cookie `json:"session_cookie"`
}

type HasRoleFromSessionParams struct {
	Environment        string       `json:"environment"`
	InfraSessionCookie *http.Cookie `json:"infra_session_cookie"`
	SessionCookie      *http.Cookie `json:"session_cookie"`
	Organization       string       `json:"organization"`
}

const (
	issuer = "briantaylorvann.com"

	guest  = "guest"
	infra  = "infra"
	client = "client"
)

func CheckGuestSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  issuer,
		Subject: guest,
	})
}

func CheckClientSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  issuer,
		Subject: client,
	})
}

func CheckInfraSession(sessionToken string) bool {
	return jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:   sessionToken,
		Issuer:  issuer,
		Subject: infra,
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
			Token:       sessionCookie.Value,
		},
		sessionCookie,
	)
	if validToken != nil {
		return true
	}
	if errValidate != nil {
		return false
	}

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
			Token:       sessionCookie.Value,
		},
		sessionCookie,
	)
	if validToken != nil {
		return true
	}
	if errValidToken != nil {
		return false
	}

	return false
}

func IsSessionValid(
	w http.ResponseWriter,
	p *IsSessionValidParams,
) bool {
	if p.InfraSessionCookie == nil {
		return false
	}

	validToken, errValidToken := fetchx.ValidateSession(
		&requests.ValidateSession{
			Environment: p.Environment,
			Token:       p.SessionCookie.Value,
		},
		p.InfraSessionCookie,
	)
	if validToken != nil {
		return true
	}
	if errValidToken != nil {
		return false
	}

	return false
}

// has role from session
func HasRoleFromSession(
	w http.ResponseWriter,
	p *HasRoleFromSessionParams,
) bool {
	if p.InfraSessionCookie == nil {
		return false
	}

	validRole, errValidRole := fetchx.ValidateRoleFromSession(
		&requests.ValidateRoleFromSession{
			Environment:  p.Environment,
			Token:        p.SessionCookie.Value,
			Organization: p.Organization,
		},
		p.InfraSessionCookie,
	)
	if errValidRole != nil {
		return false
	}
	if validRole != nil {
		return true
	}

	return false
}

// validate user
func ValidateUser(
	w http.ResponseWriter,
	p *ValidateUserParams,
) bool {
	if p.InfraSessionCookie == nil {
		return false
	}
	validRole, errValidRole := fetchx.ValidateUser(
		&requests.ValidateUser{
			Environment: p.Environment,
			Email:       p.Email,
			Password:    p.Password,
		},
		p.InfraSessionCookie,
	)
	if validRole != nil {
		return true
	}
	if errValidRole != nil {
		return false
	}

	return false
}
