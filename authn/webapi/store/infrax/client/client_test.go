package client

import (
	"testing"

	"github.com/taylor-vann/tvgtb/jwtx"
)

func TestInit(t *testing.T) {
	if Client == nil {
		t.Error("cookie jar is nil")
		return
	}
	if Client.Jar == nil {
		t.Error("cookie jar is nil")
	}

	for _, cookie := range Client.Jar.Cookies(ParsedDomain) {
		if cookie.Name == SessionCookieHeader {
			details, errDetails := jwtx.RetrieveTokenDetailsFromString(cookie.Value)
			if errDetails != nil {
				t.Error(errDetails)
			}
			if details.Payload.Sub != "infra" {
				t.Error("token subject should be infra")
			}
			t.Error(details)
		}
	}
}

func TestGuestSession(t *testing.T) {
	session, errSession := GuestSession()
	if errSession != nil {
		t.Error(errSession)
	}
	if session == "" {
		t.Error(errSession)
	}
}

func TestInfraSession(t *testing.T) {
	session, errSession := InfraSession()
	if errSession != nil {
		t.Error("session is nil")
		t.Error(errSession)
	}
	if session == "" {
		t.Error(session)
	}
}
