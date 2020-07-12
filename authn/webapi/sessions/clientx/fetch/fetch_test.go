package fetch

import (
	"net/http"
	"net/url"
	"os"
	"testing"

	"webapi/sessions/clientx/fetch/requests"
)

var Environment = os.Getenv("STAGE")

var infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
var infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")


var parsedDomain, errPasedDomain = url.Parse("https://briantaylorvann.com")
var GuestSessionTest *string
var InfraSessionTest *string

func TestGuestSession(t *testing.T) {
	session, errSession := CreateGuestSession(
		&requests.GuestSession{
			Environment: Environment,
		},
	)
	if errSession != nil {
		t.Error(errSession)
	}
	if session == nil {
		t.Error("nil response returned")
	}

	GuestSessionTest = session
}

func TestValidateGuestSession(t *testing.T) {
	if GuestSessionTest == nil {
		t.Error("guest session is nil")
	}
	resp, errResp := ValidateGuestSession(
		&requests.ValidateSession{
			Environment: Environment,
			Token: *GuestSessionTest,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
	if *resp == "" {
		t.Error("nil session returned")
	}
}

func TestValidateGuestUser(t *testing.T) {
	resp, errResp := ValidateGuestUser(
		&requests.ValidateGuestUser{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}

func TestValidateInfraRole(t *testing.T) {
	resp, errResp := ValidateInfraRole(
		&requests.ValidateInfraRole{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}

func TestCreateInfraSession(t *testing.T) {
	resp, errResp := CreateInfraSession(
		&requests.InfraSession{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}

// func TestValidateInfraSession(t *testing.T) {
// 	resp, errResp := ValidateGuestSession(
// 		requests.ValidateSession{
// 			Environment: Environment,
// 			Token: *GuestSessionTest,
// 		},
// 		&http.Cookie{
// 			Name: "briantaylorvann.com_session",
// 			Value: *GuestSessionTest,
// 		},
// 	)
// 	if errResp != nil {
// 		t.Error(errResp)
// 	}
// 	if resp == nil {
// 		t.Error("nil response returned")
// 	}
// 	if *resp == "" {
// 		t.Error("nil session returned")
// 	}
// }