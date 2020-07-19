package verifyx

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"webapi/sessions/infraclientx/sessionx"
)

var (
	Environment = os.Getenv("STAGE")
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

	GuestSessionTestCookie = &http.Cookie{
		Name: "briantaylorvann.com_session",
		Value: *session,
	}
}

func TestCreateClientxSession(t *testing.T) {
	session, errInfraSession := sessionx.Setup()
	if errInfraSession != nil {
		t.Error(errInfraSession)
	}
	if session == nil {
		t.Error("infra session is nil!")
		return
	}

	InfraSessionTestCookie = &http.Cookie{
		Name: "briantaylorvann.com_session",
		Value: *session,
	}
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
		Environment,
		InfraSessionTestCookie,
		&GuestSessionTestCookie.Value,
	) {
		t.Error("session could not be verified")
	}
}