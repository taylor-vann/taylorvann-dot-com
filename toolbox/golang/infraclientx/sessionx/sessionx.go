package sessionx

import (
	"errors"
	"net/http"
	"os"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx/requests"
)

var (
	Environment = os.Getenv("STAGE")

	infraOverlordEmail    = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")

	guestSessionRequestParams = requests.GuestSession{
		Environment: Environment,
	}
	infraSessionRequestParams = requests.InfraSession{
		Environment: Environment,
		Email:       infraOverlordEmail,
		Password:    infraOverlordPassword,
	}
)

var (
	errNilGuestSession  = errors.New("nil guest session returned")
	errNileInfraSession = errors.New("nil infra session returned")
)

var (
	GuestSession *http.Cookie
	InfraSession *http.Cookie
)

func CreateGuestSession() (*http.Cookie, error) {
	return fetchx.CreateGuestSession(
		&guestSessionRequestParams,
	)
}

func CreateInfraSession(guestSessionCookie *http.Cookie) (*http.Cookie, error) {
	return fetchx.CreateInfraSession(
		&infraSessionRequestParams,
		guestSessionCookie,
	)
}

func Setup() (*http.Cookie, error) {
	requestedGuestSession, errRequestedGuestSession := CreateGuestSession()
	if errRequestedGuestSession != nil {
		return nil, errRequestedGuestSession
	}
	if requestedGuestSession == nil {
		return nil, errNilGuestSession
	}

	requestedInfraSession, errRequestedInfraSession := CreateInfraSession(
		requestedGuestSession,
	)
	if errRequestedInfraSession != nil {
		return nil, errRequestedInfraSession
	}
	if requestedInfraSession == nil {
		return nil, errNileInfraSession
	}

	GuestSession = requestedGuestSession
	InfraSession = requestedInfraSession

	return requestedInfraSession, nil
}
