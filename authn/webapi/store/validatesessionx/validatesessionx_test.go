package validatesessionx

import (
	"testing"
)

func TestFetchGuest(t *testing.T) {
	resp, errResp := FetchGuestSession()

	if errResp != nil {
		t.Error(errResp)
	}
 
	if resp == "" {
		t.Error("empty session string returned")
	}
}

func TestValidateGuestSession(t *testing.T) {
	resp, errResp := FetchGuestSession()
	if errResp != nil {
		t.Error(errResp)
	}
 
	if resp == "" {
		t.Error("empty session string returned")
	}

	respValidate, errRespValidagte := ValidateGuestSession(resp)

	if errResp != nil {
		t.Error(errRespValidagte)
	}
 
	if respValidate == false {
		t.Error("failed to validate session token")
	}
}