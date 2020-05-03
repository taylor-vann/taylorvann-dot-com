package hooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"webapi/interfaces/jwtx"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
)

func TestCreateGuestSessionBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp.Error())
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
	requestBody := requests.Body{
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
		t.Error(errReq.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status == http.StatusOK {
		t.Error("handler returned incorrect status code, should not be 200")
	}
}

func TestCreateGuestSession(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.SessionParams{
			Environment: "LOCAL",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestCreateDocumentSession(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateDocumentSession,
		Params: requests.SessionParams{
			Environment: "LOCAL",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}

	if httpTest.Code != http.StatusOK {
		t.Error(httpTest.Code)
	}
}

func TestCreateResetPasswordSessionBadRequest(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateUpdatePasswordSession,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreateCreateAccountSession(t *testing.T) {
	email := "something@darkside.complete"
	// public session from guest sesion
	requestBodyPublic := requests.Body{
		Action: CreateCreateAccountSession,
		Params: requests.AccountParams{
			Environment: "LOCAL",
			Email: email,
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
		t.Error(errReq.Error())
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreateUpdatePasswordSession(t *testing.T) {
	email := "something@darkside.complete"
	// public session from guest sesion
	requestBodyPublic := requests.Body{
		Action: CreateUpdatePasswordSession,
		Params: requests.AccountParams{
			Environment: "LOCAL",
			Email: email,
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
		t.Error(errReq.Error())
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)

	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestCreateUpdateEmailSession(t *testing.T) {
	email := "something@darkside.complete"
	// public session from guest sesion
	requestBodyPublic := requests.Body{
		Action: CreateUpdateEmailSession,
		Params: &requests.AccountParams{
			Environment: "LOCAL",
			Email: email,
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
		t.Error(errReq.Error())
	}

	httpTestPublic := httptest.NewRecorder()
	handlerPublic := http.HandlerFunc(Mutation)
	handlerPublic.ServeHTTP(httpTestPublic, req)

	statusPublic := httpTestPublic.Code
	if statusPublic != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}
}

func TestUpdateSession(t *testing.T) {
	sessionRequestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.SessionParams{
			Environment: "LOCAL",
		},
	}

	sessionMarshalBytes := new(bytes.Buffer)
	json.NewEncoder(sessionMarshalBytes).Encode(sessionRequestBody)
	respSession, errSessionResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		sessionMarshalBytes,
	)
	if errSessionResp != nil {
		t.Error(errSessionResp)
	}

	httpTestSession := httptest.NewRecorder()
	handlerSession := http.HandlerFunc(Mutation)
	handlerSession.ServeHTTP(httpTestSession, respSession)
	if respSession.Body == nil {
		t.Error("response body is nil")
	}
	var responseBodySession responses.Body
	errJSONSession := json.NewDecoder(httpTestSession.Body).Decode(&responseBodySession)
	if errJSONSession != nil {
		t.Error(errJSONSession)
	}
	if responseBodySession.Session == nil {
		t.Error("nil session returned")
	}

	time.Sleep(10 * time.Millisecond)

	// public session from guest sesion
	requestBodyPublic := requests.Body{
		Action: UpdateSession,
		Params: &requests.Update{
			Environment: "LOCAL",
			SessionToken: responseBodySession.Session.SessionToken,
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
		t.Error(errReq.Error())
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
	requestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.SessionParams{
			Environment: "LOCAL",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
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
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}
	// public session from guest sesion
	requestBodyValidate := requests.Body{
		Action: ValidateSession,
		Params: &requests.Read{
			Environment: "LOCAL",
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
		t.Error(errReq.Error())
	}

	httpTestValidate := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Query)

	handlerRemove.ServeHTTP(httpTestValidate, req)

	statusRemove := httpTestValidate.Code
	if statusRemove != http.StatusOK {
		t.Error("validate handler returned incorrect status code")
	}
}

func TestDeleteSession(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.SessionParams{
			Environment: "LOCAL",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
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
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}

	// get signature
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		responseBody.Session.SessionToken,
	)
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}
	// public session from guest sesion
	requestBodyRemove := requests.Body{
		Action: DeleteSession,
		Params: &requests.Delete{
			Environment: "LOCAL",
			Signature: tokenDetails.Signature,
		},
	}

	marshalBytesRemove := new(bytes.Buffer)
	errRemove := json.NewEncoder(marshalBytesRemove).Encode(requestBodyRemove)
	if errRemove != nil {
		t.Error(errRemove.Error())
	}
	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		marshalBytesRemove,
	)
	if errReq != nil {
		t.Error(errReq.Error())
	}

	httpTestRemove := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Mutation)

	handlerRemove.ServeHTTP(httpTestRemove, req)

	statusRemove := httpTestRemove.Code
	if statusRemove != http.StatusOK {
		t.Error("remove handler returned incorrect status code")
	}
}
