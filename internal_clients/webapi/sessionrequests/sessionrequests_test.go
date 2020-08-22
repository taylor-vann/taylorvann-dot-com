package sessionrequests

import (
	"net/http"
	"testing"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
)


var (
	GuestSessionTestCookie *http.Cookie
	ClientSessionTestCookie *http.Cookie
)

// guest session
func TestCreateGuestSession(t *testing.T) {
	session, errInfraSession := sessionx.CreateGuestSession()
	if errInfraSession != nil {
		t.Error(errInfraSession)
		return
	}
	if session == nil {
		t.Error("guest session is nil!")
		return
	}

	// set for verification on next text
	GuestSessionTestCookie = session
}

// clientx session
func TestCreateClientxSession(t *testing.T) {
	session, errInfraSession := sessionx.CreateInfraSession(GuestSessionTestCookie)
	if errInfraSession != nil {
		t.Error(errInfraSession)
	}
	if session == nil {
		t.Error("infra session is nil!")
		return
	}

	// set for verification on next text
	ClientSessionTestCookie = session
}
