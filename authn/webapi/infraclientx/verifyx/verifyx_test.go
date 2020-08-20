package verifyx

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"webapi/infraclientx/sessionx"
)

var (
	Environment = os.Getenv("STAGE")

	infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")
)

var (
	GuestSessionTestCookie *http.Cookie
	InfraSessionTestCookie *http.Cookie
)

func TestCreateGuestSession(t *testing.T) {
	session, errInfraSession := sessionx.CreateGuestSession()
	if errInfraSession != nil {
		t.Error(errInfraSession)
	}
	if session == nil {
		t.Error("infra session is nil!")
		return
	}

	GuestSessionTestCookie = session
}

func TestCreateClientxSession(t *testing.T) {
	session, errInfraSession := sessionx.CreateInfraSession(GuestSessionTestCookie)
	if errInfraSession != nil {
		t.Error(errInfraSession)
	}
	if session == nil {
		t.Error("infra session is nil!")
		return
	}

	InfraSessionTestCookie = session
}

func TestCheckGuestSession(t *testing.T) {
	if GuestSessionTestCookie == nil {
		t.Error("guest session is nil")
	}
	if !CheckGuestSession(GuestSessionTestCookie.Value) {
		t.Error("guest session is not valid")
	}
}

func TestCheckInfraSession(t *testing.T) {
	if InfraSessionTestCookie == nil {
		t.Error("guest session is nil")
	}
	if !CheckInfraSession(InfraSessionTestCookie.Value) {
		t.Error("guest session is not valid")
	}
}

func TestIsGuestSessionValid(t *testing.T) {
	if GuestSessionTestCookie == nil {
		t.Error("guest session is nil")
	}

	htr := httptest.NewRecorder()

	if !IsGuestSessionValid(htr, Environment, GuestSessionTestCookie) {
		t.Error("guest session is not valid")
	}
}

func TestIsInfraSessionValid(t *testing.T) {
	if InfraSessionTestCookie == nil {
		t.Error("guest session is nil")
	}

	htr := httptest.NewRecorder()

	if !IsInfraSessionValid(htr, Environment, InfraSessionTestCookie) {
		t.Error("guest session is not valid")
	}
}

func TestIsSessionValid(t *testing.T) {
	if InfraSessionTestCookie == nil {
		t.Error("guest session is nil")
	}

	htr := httptest.NewRecorder()

	if !IsSessionValid(
		htr,
		&IsSessionValidParams{
			Environment: Environment,
			InfraSessionCookie: InfraSessionTestCookie,
			SessionCookie: GuestSessionTestCookie,
		},
	) {
		t.Error("session could not be verified")
	}
}

func TestHasRoleFromSession(t *testing.T) {
	if InfraSessionTestCookie == nil {
		t.Error("guest session is nil")
	}

	htr := httptest.NewRecorder()

	if !HasRoleFromSession(
		htr,
		&HasRoleFromSessionParams{
			Environment: Environment,
			InfraSessionCookie: InfraSessionTestCookie,
			SessionCookie: InfraSessionTestCookie,
			Organization: "AUTHN_ADMIN",
		},
	) {
		t.Error("session could not be verified")
	}
}

func TestValidateUser(t *testing.T) {
	if InfraSessionTestCookie == nil {
		t.Error("guest session is nil")
	}

	htr := httptest.NewRecorder()

	if !ValidateUser(
		htr,
		&ValidateUserParams{
			Environment: Environment,
			InfraSessionCookie: InfraSessionTestCookie,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
	) {
		t.Error("session could not be verified")
	}
}