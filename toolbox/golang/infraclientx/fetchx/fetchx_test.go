package fetchx

import (
	"net/http"
	"os"
	"testing"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/fetchx/requests"
)

var (
	Environment = os.Getenv("STAGE")

	infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")
)

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

	InfraSessionTest = resp
}

func TestValidateSession(t *testing.T) {
	if GuestSessionTest == nil {
		t.Error("guest session is nil")
	}
	if InfraSessionTest == nil {
		t.Error("guest session is nil")
	}
	resp, errResp := ValidateSession(
		&requests.ValidateSession{
			Environment: Environment,
			Token: *GuestSessionTest,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *InfraSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}

func TestValidateRoleFromSession(t *testing.T) {
	if GuestSessionTest == nil {
		t.Error("guest session is nil")
	}
	if InfraSessionTest == nil {
		t.Error("guest session is nil")
	}
	resp, errResp := ValidateRoleFromSession(
		&requests.ValidateRoleFromSession{
			Environment: Environment,
			Token: *InfraSessionTest,
			Organization: "AUTHN_ADMIN",
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *InfraSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}

func TestValidateUser(t *testing.T) {
	if GuestSessionTest == nil {
		t.Error("guest session is nil")
	}
	if InfraSessionTest == nil {
		t.Error("guest session is nil")
	}
	resp, errResp := ValidateUser(
		&requests.ValidateUser{
			Environment: Environment,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
		&http.Cookie{
			Name: "briantaylorvann.com_session",
			Value: *InfraSessionTest,
		},
	)
	if errResp != nil {
		t.Error(errResp)
	}
	if resp == nil {
		t.Error("nil response returned")
	}
}