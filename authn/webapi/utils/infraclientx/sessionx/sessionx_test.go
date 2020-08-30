package sessionx

import (
	"testing"
)

func TestGuestSession(t *testing.T) {
	session, errSession := CreateGuestSession()
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil guest session returned")
	}
}

func TestInfraSession(t *testing.T) {
	session, errSession := CreateGuestSession()
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil guest session returned")
	}
	infraSession, errInfraSession := CreateInfraSession(session)
	if errInfraSession != nil {
		t.Error(errSession)
	}
	if infraSession == nil {
		t.Error("nil infra session returned")
	}
}

func TestSetup(t *testing.T) {
	session, errSession := Setup()
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil infra session returned")
	}
	if GuestSession == nil {
		t.Error("guest session is nil")
	}
	if InfraSession == nil {
		t.Error("infra session is nil")
	}
}
