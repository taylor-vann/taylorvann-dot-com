package sessions

import (
	"bytes"
	"encoding/json"
	// "fmt"

	"net/http"
	"net/http/httptest"
	"testing"
	"webapi/hooks/constants"
	// "webapi/interfaces/jwtx"
	// "webapi/sessions"
)

func TestCreateGuestSessionBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		nil,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestCreateGuestSessionBadHeadersRequest(t *testing.T) {
	requestBody := RequestBodyParams{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	req, errReq := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytes,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	req.Header.Set(constants.SessionTokenHeader, "asdfjkl;")
	req.Header.Set(constants.CsrfTokenHeader, "asdfjkl;")

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestCreateGuestSession(t *testing.T) {
	requestBody := RequestBodyParams{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreatePublicPasswordResetSessionBadHeaders(t *testing.T) {
	requestBody := RequestBodyParams{
		Action: CreatePublicPasswordResetSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreatePublicPasswordResetSession(t *testing.T) {
	requestBody := RequestBodyParams{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}

	// public session from guest sesion
	requestBodyPublic := RequestBodyParams{
		Action: CreatePublicPasswordResetSession,
	}

	marshalBytesPublic := new(bytes.Buffer)
	json.NewEncoder(marshalBytesPublic).Encode(requestBodyPublic)
	req, errReq := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytesPublic,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	req.Header.Set(
		constants.SessionTokenHeader,
		httpTest.Header().Get(constants.SessionTokenHeader),
	)

	req.Header.Set(
		constants.CsrfTokenHeader,
		httpTest.Header().Get(constants.CsrfTokenHeader),
	)
	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestRemoveSession(t *testing.T) {
	requestBody := RequestBodyParams{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("guest handler returned incorrect status code")
	}

	// public session from guest sesion
	requestBodyRemove := RequestBodyParams{
		Action: RemoveSession,
	}

	marshalBytesRemove := new(bytes.Buffer)
	json.NewEncoder(marshalBytesRemove).Encode(requestBodyRemove)
	req, errReq := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytesRemove,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestRemove := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Mutation)

	req.Header.Set(
		constants.SessionTokenHeader,
		httpTest.Header().Get(constants.SessionTokenHeader),
	)

	req.Header.Set(
		constants.CsrfTokenHeader,
		httpTest.Header().Get(constants.CsrfTokenHeader),
	)
	handlerRemove.ServeHTTP(httpTestRemove, req)

	statusRemove := httpTestRemove.Code
	if statusRemove != http.StatusOK {
		t.Error("remove handler returned incorrect status code")
	}

	t.Error("Fail because it's new")
}
