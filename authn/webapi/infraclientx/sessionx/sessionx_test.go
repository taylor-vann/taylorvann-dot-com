package sessionx

import (
	"testing"
	"net/http"
)

func TestGuestSession(t *testing.T) {
	session, errSession := CreateGuestSession()
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil guest session returned")
	}
	if GuestSession == nil {
		t.Error("guest session is nil")
	}
}

func TestInfraSession(t *testing.T) {
	if GuestSession == nil {
		t.Error("GuestSession is nil")
		return
	}
	session, errSession := CreateInfraSession(&http.Cookie{
		Name: "briantaylorvann.com_session",
		Value: *GuestSession,
	})
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil infra session returned")
	}
	if InfraSession == nil {
		t.Error("infra session is nil")
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
