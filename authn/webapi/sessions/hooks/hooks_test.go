package hooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"os"

	"log"

	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"

	"webapi/sessions/clientx/sessionx"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

var (
	InfraEmail 		= os.Getenv("INFRA_OVERLORD_EMAIL")
	InfraPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")
)

var (
	GuestSessionTest string
	ClientSessionTest *http.Cookie
)

func TestCreateGuestSession(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.Guest{
			Environment: "LOCAL",
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
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
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, resp)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}

	details, errDetails := jwtx.RetrieveTokenDetailsFromString(responseBody.Session.Token)
	if errDetails != nil {
		t.Error(errDetails)
		return
	}
	if details.Payload.Iss != "briantaylorvann.com" {
		t.Error(details.Payload.Iss)
	}

	// set for verification on next text
	GuestSessionTest = responseBody.Session.Token
}

// clientx session
func TestCreateClientxSession(t *testing.T) {
	log.Println("TestCreateClientxSession")
	cookie, errInfraSession := sessionx.Setup()
	if errInfraSession != nil {
		t.Error(errInfraSession)
	}
	if cookie == nil {
		t.Error("infra session is nil!")
	}

	// set for verification on next text
	ClientSessionTest = cookie
}

func TestCreateGuestSessionBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/m/sessions/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, resp)

	if htr.Code != http.StatusBadRequest {
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

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, req)

	status := htr.Code
	if status == http.StatusOK {
		t.Error("handler returned incorrect status code, should not be 200")
	}
}

func TestValidateGuestSession(t *testing.T) {
	requestBody := requests.Body{
		Action: ValidateGuestSession,
		Params: requests.Validate{
			Environment: "LOCAL",
			Token: GuestSessionTest,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	resp, errResp := http.NewRequest(
		"POST",
		"/q/sessions/",
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
	resp.AddCookie(&http.Cookie{
		Name: "briantaylorvann.com_session",
		Value: GuestSessionTest,
	})
	
	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)
	handler.ServeHTTP(htr, resp)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestCreateClientSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(ClientSessionTest.Value)
	if errDetails != nil {
		t.Error(errDetails)
		return
	}

	audAsInt64, errAudAsInt64 := strconv.ParseInt(details.Payload.Aud, 10, 64)
	if errAudAsInt64 != nil {
		t.Error(errAudAsInt64)
		return
	}

	requestBody := requests.Body{
		Action: CreateClientSession,
		Params: requests.User{
			Environment: "LOCAL",
			UserID: audAsInt64,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		bytes.NewBuffer(marshalBody),
	)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	if req == nil {
		t.Error(req)
		return
	}
	req.AddCookie(ClientSessionTest)

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestCreateAccountSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(ClientSessionTest.Value)
	if errDetails != nil {
		t.Error(errDetails)
		return
	}

	audAsInt64, errAudAsInt64 := strconv.ParseInt(details.Payload.Aud, 10, 64)
	if errAudAsInt64 != nil {
		t.Error(errAudAsInt64)
		return
	}

	requestBody := requests.Body{
		Action: CreateCreateAccountSession,
		Params: requests.User{
			Environment: "LOCAL",
			UserID: audAsInt64,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		bytes.NewBuffer(marshalBody),
	)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	if req == nil {
		t.Error(req)
		return
	}
	req.AddCookie(ClientSessionTest)

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestCreateUpdatePasswordSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(ClientSessionTest.Value)
	if errDetails != nil {
		t.Error(errDetails)
		return
	}

	audAsInt64, errAudAsInt64 := strconv.ParseInt(details.Payload.Aud, 10, 64)
	if errAudAsInt64 != nil {
		t.Error(errAudAsInt64)
		return
	}

	requestBody := requests.Body{
		Action: CreateUpdatePasswordSession,
		Params: requests.User{
			Environment: "LOCAL",
			UserID: audAsInt64,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		bytes.NewBuffer(marshalBody),
	)
	req.AddCookie(ClientSessionTest)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	if req == nil {
		t.Error(req)
		return
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestCreateUpdateEmailSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(ClientSessionTest.Value)
	if errDetails != nil {
		t.Error(errDetails)
		return
	}

	audAsInt64, errAudAsInt64 := strconv.ParseInt(details.Payload.Aud, 10, 64)
	if errAudAsInt64 != nil {
		t.Error(errAudAsInt64)
		return
	}

	requestBody := requests.Body{
		Action: CreateUpdateEmailSession,
		Params: requests.User{
			Environment: "LOCAL",
			UserID: audAsInt64,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		bytes.NewBuffer(marshalBody),
	)
	req.AddCookie(ClientSessionTest)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	if req == nil {
		t.Error(req)
		return
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestValidateSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	requestBody := requests.Body{
		Action: ValidateSession,
		Params: requests.Validate{
			Environment: "LOCAL",
			Token: ClientSessionTest.Value,
		},
	}

	marshalBody, errMarshalBody := json.Marshal(requestBody)
	if errMarshalBody != nil {
		t.Error(errMarshalBody)
		return
	}

	req, errReq := http.NewRequest(
		"POST",
		"/m/sessions/",
		bytes.NewBuffer(marshalBody),
	)
	req.AddCookie(ClientSessionTest)
	if errReq != nil {
		t.Error(errReq.Error())
		return
	}
	if req == nil {
		t.Error(req)
		return
	}

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)
	handler.ServeHTTP(htr, req)

	if htr.Code != http.StatusOK {
		t.Error(htr.Code)
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Session == nil {
		t.Error("nil session returned")
	}
}

func TestDeleteSession(t *testing.T) {
	if ClientSessionTest == nil {
		t.Error("client session is nil")
		return
	}
	requestBody := requests.Body{
		Action: CreateGuestSession,
		Params: requests.Guest{
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

	htr := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(htr, resp)

	status := htr.Code
	if status != http.StatusOK {
		t.Error("guest handler returned incorrect status code")
		return
	}

	var responseBody responses.Body
	errJSON := json.NewDecoder(htr.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
		return
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		responseBody.Session.Token,
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
	req.AddCookie(ClientSessionTest)
	if errReq != nil {
		t.Error(errReq.Error())
	}

	htrRemove := httptest.NewRecorder()
	handlerRemove := http.HandlerFunc(Mutation)
	handlerRemove.ServeHTTP(htrRemove, req)

	statusRemove := htrRemove.Code
	if statusRemove != http.StatusOK {
		t.Error("remove handler returned incorrect status code")
	}
}
