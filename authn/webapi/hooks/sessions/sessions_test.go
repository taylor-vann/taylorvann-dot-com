package sessions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"webapi/hooks/sessions/errors"
)

func TestCreateGuestSessionBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
	requestBody := MutationRequestBody{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("handler returned incorrect status code, should be 200")
	}
}

func TestCreateGuestSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp)
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody MutationResponseBody
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON)
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}

	if httpTest.Code != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreateDocumentSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateDocumentSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp)
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody MutationResponseBody
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON)
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}

	if httpTest.Code != http.StatusOK {
		t.Error(httpTest.Code)
	}
}

func TestCreateResetPasswordSessionBadRequest(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateResetPasswordSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp)
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreateResetPasswordSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateDocumentSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
		t.Error("guest session returned incorrect status code")
		return
	}

	time.Sleep(10 * time.Millisecond)

	// decode body to response
	var responseBody MutationResponseBody
	errResponseBody := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errResponseBody != nil {
		t.Error(errResponseBody)
		return
	}

	email := "something@darkside.complete"
	// public session from guest sesion
	requestBodyPublic := MutationRequestBody{
		Action: CreateResetPasswordSession,
		Params: &MutationRequestPayload{
			SessionToken: responseBody.Session.SessionToken,
			Credentials: &errors.Credentials{
				Email: &email,
			},
		},
	}

	marshalBytesPublic := new(bytes.Buffer)
	json.NewEncoder(marshalBytesPublic).Encode(requestBodyPublic)
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytesPublic,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		fmt.Println(httpTestPublic)
		t.Error("handler returned incorrect status code")
	}
}


func TestCreateUpdateEmailSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateDocumentSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
		t.Error("guest session returned incorrect status code")
		return
	}

	time.Sleep(10 * time.Millisecond)

	// decode body to response
	var responseBody MutationResponseBody
	errResponseBody := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errResponseBody != nil {
		t.Error(errResponseBody)
		return
	}

	email := "something@darkside.complete"
	// public session from guest sesion
	requestBodyPublic := MutationRequestBody{
		Action: CreateUpdateEmailSession,
		Params: &MutationRequestPayload{
			SessionToken: responseBody.Session.SessionToken,
			Credentials: &errors.Credentials{
				Email: &email,
			},
		},
	}

	marshalBytesPublic := new(bytes.Buffer)
	json.NewEncoder(marshalBytesPublic).Encode(requestBodyPublic)
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytesPublic,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		fmt.Println(httpTestPublic)
		t.Error("handler returned incorrect status code")
	}
}

func TestUpdateSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
		return
	}

	// decode body to response
	var responseBody MutationResponseBody
	errResponseBody := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errResponseBody != nil {
		t.Error(errResponseBody)
		return
	}

	time.Sleep(10 * time.Millisecond)

	// public session from guest sesion
	requestBodyPublic := MutationRequestBody{
		Action: UpdateSession,
		Params: &MutationRequestPayload{
			SessionToken: responseBody.Session.SessionToken,
		},
	}

	marshalBytesPublic := new(bytes.Buffer)
	json.NewEncoder(marshalBytesPublic).Encode(requestBodyPublic)
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytesPublic,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestValidateSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
		return
	}

	// this is the new stuffs
	var responseBody MutationResponseBody
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON)
		return
	}
	// public session from guest sesion
	requestBodyValidate := QueryRequestBody{
		Action: ValidateSession,
		Params: &QueryRequestPayload{
			SessionToken: responseBody.Session.SessionToken,
		},
	}

	marshalBytesValidate := new(bytes.Buffer)
	errRemove := json.NewEncoder(marshalBytesValidate).Encode(requestBodyValidate)
	if errRemove != nil {
		t.Error(errRemove)
	}
	req, errReq := http.NewRequest(
		"POST",
		"/sessions/q/",
		marshalBytesValidate,
	)
	if errReq != nil {
		t.Error("error making validating request")
	}

	httpTestValidate := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Query)

	handlerRemove.ServeHTTP(httpTestValidate, req)

	statusRemove := httpTestValidate.Code
	if statusRemove != http.StatusOK {
		t.Error("validate handler returned incorrect status code")
	}
}

func TestRemoveSession(t *testing.T) {
	requestBody := MutationRequestBody{
		Action: CreateGuestSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
		return
	}

	// this is the new stuffs
	var responseBody MutationResponseBody
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON)
		return
	}
	// public session from guest sesion
	requestBodyRemove := MutationRequestBody{
		Action: RemoveSession,
		Params: &MutationRequestPayload{
			SessionToken: responseBody.Session.SessionToken,
		},
	}

	marshalBytesRemove := new(bytes.Buffer)
	errRemove := json.NewEncoder(marshalBytesRemove).Encode(requestBodyRemove)
	if errRemove != nil {
		t.Error(errRemove)
	}
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytesRemove,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTestRemove := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Mutation)

	handlerRemove.ServeHTTP(httpTestRemove, req)

	statusRemove := httpTestRemove.Code
	if statusRemove != http.StatusOK {
		t.Error("remove handler returned incorrect status code")
	}
}
