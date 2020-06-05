package sessionx

import (
	"testing"

	"github.com/taylor-vann/tvgtb/jwtx"
)

func TestInit(t *testing.T) {
	if client == nil {
		t.Error("cookie jar is nil")
		return
	}
	if client.Jar == nil {
		t.Error("cookie jar is nil")
	}

	for _, cookie := range client.Jar.Cookies(parsedDomain) {
		if cookie.Name == sessionCookieHeader {
			details, errDetails := jwtx.RetrieveTokenDetailsFromString(cookie.Value)
			if errDetails != nil {
				t.Error(errDetails)
			}
			if details.Payload.Sub != "infra" {
				t.Error("token subject should be infra")
			}
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
	session, errSession := infraSession()
	if errSession != nil {
		t.Error("session is nil")
		t.Error(errSession)
	}
	if session == "" {
		t.Error(session)
	}
}
