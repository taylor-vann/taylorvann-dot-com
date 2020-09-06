package sessions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"webapi/sessions/requests"
	"webapi/sessions/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
)

var (
	InfraEmail    = os.Getenv("INFRA_OVERLORD_EMAIL")
	InfraPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")
)

var (
	GuestSessionTestCookie  *http.Cookie
	ClientSessionTestCookie *http.Cookie
)

func TestSetupSessionX(t *testing.T) {
	sessionx.Setup()
	// set for verification on next text
	GuestSessionTestCookie = sessionx.GuestSession
	ClientSessionTestCookie = sessionx.InfraSession
}

func TestGuestSession(t *testing.T) {
	requestBody := requests.CreateGuestSessionParams{
		Environment: "DEVELOPMENT",
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	resp, errResp := http.NewRequest(
		"POST",
		"/request_guest_session/",
		bytes.NewBuffer(marshalBody),
	)
	if resp == nil {
		t.Error(resp)
		return
	}
	if errResp != nil {
		t.Error(errResp.Error())
		return
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestGuestSession)
	handler.ServeHTTP(htr, resp)

	if htr.Code != http.StatusOK {
		t.Error("incorrect status code")
		t.Error(htr.Code)
	}

	result := htr.Result()

	var responseBody responses.Body
	errJSON := json.NewDecoder(result.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}
	if responseBody.Errors != nil {
		t.Error("errors returned")
		return
	}
}

func TestRequestClientSession(t *testing.T) {
	requestBody := requests.CreateClientSessionParams{
		Environment: "DEVELOPMENT",
		Email:       InfraEmail,
		Password:    InfraPassword,
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	resp, errResp := http.NewRequest(
		"POST",
		"/request_client_session/",
		bytes.NewBuffer(marshalBody),
	)
	if resp == nil {
		t.Error(resp)
		return
	}
	if errResp != nil {
		t.Error(errResp.Error())
		return
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(RequestClientSession)
	handler.ServeHTTP(htr, resp)

	if htr.Code != http.StatusOK {
		t.Error("incorrect status code")
		t.Error(htr.Code)
	}

	result := htr.Result()

	var responseBody responses.Body
	errJSON := json.NewDecoder(result.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}
	if responseBody.Errors != nil {
		t.Error("errors returned")
		return
	}
}

func TestRemoveSession(t *testing.T) {
	if ClientSessionTestCookie == nil {
		t.Error("client session test cookie is nil")
		return
	}

	requestBody := requests.RemoveSessionParams{
		Environment: "DEVELOPMENT",
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/remove_session/",
		bytes.NewBuffer(marshalBody),
	)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	req.AddCookie(ClientSessionTestCookie)

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveSession)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error("incorrect status code")
		t.Error(htr.Code)
	}

	result := htr.Result()

	var responseBody responses.Body
	errJSON := json.NewDecoder(result.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}
	if responseBody.Errors != nil {
		t.Error("errors returned")
		return
	}
}
