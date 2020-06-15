package fetch

import (
	"net/http"
	"net/url"
	"testing"

	"os"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/requests"

)

var Environment = os.Getenv("STAGE")

var infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
var infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")


var parsedDomain, errPasedDomain = url.Parse("https://briantaylorvann.com")
var GuestSessionTest string
var InfraSessionTest string

func TestGuestSession(t *testing.T) {
	resp, errResp := sessionx.GuestSession()
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == "" {
		t.Error("nil response returned")
	}

	GuestSessionTest = resp
}

func TestValidateGuestSession(t *testing.T) {
	resp, errResp := ValidateGuestSession(
		requests.ValidateGuestSession{
			Environment: Environment,
			
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == "" {
		t.Error("nil response returned")
	}
}

func TestValidateUser(t *testing.T) {
	resp, errResp := ValidateGuestUser(
		requests.ValidateGuestUser{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: GuestSessionTest,
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
		requests.ValidateGuestUser{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: GuestSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}