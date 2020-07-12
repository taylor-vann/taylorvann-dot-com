package clientx

import (
	"testing"

	"webapi/sessions/clientx/fetch/requests"
	"webapi/sessions/clientx/sessionx"
)

func TestDo(t *testing.T) {
	if sessionx.Session == nil {
		t.Error("session is nil")
		return
	}

	params := requests.ValidateSession{
		Environment: sessionx.Environment,
		Token: sessionx.Session.Value,
	}
	resp, errResp := ValidateSession(params)
	if errResp != nil {
		t.Error(errResp)
		return
	}

	if resp != sessionx.Session.Value {
		t.Error("session is not valid!")
	}
}
