package infrasessionx

import (
	"errors"
	"net/http"
	"os"

	"webapi/sessions/clientx/infrafetchx"
	"webapi/sessions/clientx/infrafetchx/requests"
)

var (
	Environment = os.Getenv("STAGE")

	infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")
	guestSessionRequestParams = requests.GuestSession{
		Environment: Environment,
	}
	infraSessionRequestParams = requests.InfraSession {
		Environment: Environment,
		Email: infraOverlordEmail,
		Password: infraOverlordPassword,
	}
)

var (
	GuestSession *string
	InfraSession *string
)

func CreateGuestSession() (*string, error) {
	session, errSession := infrafetchx.CreateGuestSession(
		&guestSessionRequestParams,
	)
	GuestSession = session

	return session, errSession
}

func CreateInfraSession(guestSessionCookie *http.Cookie) (*string, error) {
	session, errSession := infrafetchx.CreateInfraSession(
		&infraSessionRequestParams,
		guestSessionCookie,
	)
	InfraSession = session

	return session, errSession
}

func Setup() (*string, error) {
	guestSession, errGuestSession := CreateGuestSession()
	if errGuestSession != nil {
		return nil, errGuestSession
	}
	if guestSession == nil {
		return nil, errors.New("nil guest session returned")
	}

	return CreateInfraSession(
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *guestSession,
		},
	)
}