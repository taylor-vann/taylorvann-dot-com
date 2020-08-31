package verifyx

import (
	"net/http"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx/requests"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx"

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
	if validToken != nil && errValidate == nil {
		return true
	}

	return false
}

func IsInfraSessionValid(
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
	if validToken != nil && errValidToken == nil {
		return true
	}

	return false
}

func IsSessionValid(p *IsSessionValidParams) bool {
	if p.InfraSessionCookie == nil {
		return false
	}
	if p.SessionCookie == nil {
		return false
	}
	if !CheckInfraSession(p.InfraSessionCookie.Value) {
		return false
	}

	validToken, errValidToken := fetchx.ValidateSession(
		&requests.ValidateSession{
			Environment: p.Environment,
			Token:       p.SessionCookie.Value,
		},
		p.InfraSessionCookie,
	)
	if validToken != nil && errValidToken == nil {
		return true
	}

	return false
}

// has role from session
func HasRoleFromSession(p *HasRoleFromSessionParams) bool {
	if p.InfraSessionCookie == nil {
		return false
	}
	if p.SessionCookie == nil {
		return false
	}
	if !CheckInfraSession(p.InfraSessionCookie.Value) {
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
	if validRole != nil && errValidRole == nil {
		return true
	}

	return false
}

// validate user
func ValidateUser(p *ValidateUserParams) bool {
	if p.InfraSessionCookie == nil {
		return false
	}

	validUser, errValidUser := fetchx.ValidateUser(
		&requests.ValidateUser{
			Environment: p.Environment,
			Email:       p.Email,
			Password:    p.Password,
		},
		p.InfraSessionCookie,
	)
	
	if validUser != nil && errValidUser == nil {
		return true
	}

	return false
}
