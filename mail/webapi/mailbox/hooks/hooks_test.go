// brian taylor vann
// taylorvann dot com

package hooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"webapi/mailbox/hooks/requests"
	"webapi/mailbox/hooks/responses"
)

var testEmail = requests.EmailParams{
	RecipientAddress:	"brian.t.vann@gmail.com",
	Subject: "taylorvann.com integration test!",
	Body: "Hey brian, it's brian!\n\nThis is an unit test. You can ignore it :)\n\nBest\nBrian",
	ReplyAddress: "unit_tests@taylorvann.com",
	ReplyName: "Integration Tests",
}

func TestSendNoReplyBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/sendonly/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(NoReply)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestSendNoReplyBadBodyRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/sendonly/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(NoReply)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestSendNoReply(t *testing.T) {
	requestBody := requests.Body{
		Action: CreateSendonlyEmail,
		Params: &testEmail,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	req, errReq := http.NewRequest(
		"POST",
		"/sendonly/",
		marshalBytes,
	)
	if errReq != nil {
		t.Error(errReq.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(NoReply)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("handler returned incorrect status code, should be 200")
		t.Error(status)
		return
	}

	var responseBodyNoReply responses.Body
	errJSONSession := json.NewDecoder(httpTest.Body).Decode(&responseBodyNoReply)
	if errJSONSession != nil {
		t.Error(errJSONSession)
	}
	if responseBodyNoReply.Mail == nil {
		t.Error("nil mail confirmation returned")
	}
	if responseBodyNoReply.Errors != nil {
		t.Error(*responseBodyNoReply.Errors.Default)
	}
}
