package mux

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"testing"
	"bytes"

	"webapi/setterx"
)

const (
	testKind = "test kind"
	testMessage = "test message"
	statusOk = 200
	statusNotOk = 400
)

func TestCreateMux(t *testing.T) {
	proxyMux := CreateMux()
	if proxyMux == nil {
		t.Fail()
		t.Logf("proxyMux was not created")
	}
}

func TestWriteError(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeError(testRecorder, testKind, testMessage)

	if testRecorder.Code != statusNotOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusNotOk, ", found: ", testRecorder.Code))
	}

	var errors ErrorDeclarations
	json.NewDecoder(testRecorder.Body).Decode(&errors)

	if len(errors) == 0 {
		t.Fail()
		t.Logf("error array has a length of zero")
		return
	}

	if errors[0].Kind != testKind {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testKind, ", found: ", errors[0].Kind))
	}

	if errors[0].Message != testMessage {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found: ", errors[0].Message))
	}
}

func TestWriteGetEntry(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeGetEntry(testRecorder, testMessage)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}

	var result string
	json.NewDecoder(testRecorder.Body).Decode(&result)

	if result == "" {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found nil"))
		return
	}

	if result != testMessage {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found: ", result))
	}
}

func TestWriteSetEntry(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeGetEntry(testRecorder, testMessage)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}
}

func TestSetEntryRequest(t *testing.T) {
	entry := setterx.SetBody{
		Address: testKind,
		Entry: testMessage,
	}

	reqJson, errReqJson := json.Marshal(entry)
	if errReqJson != nil {
		t.Fail()
		t.Logf(errReqJson.Error())
		return
	}

	req, errReq := http.NewRequest("GET", "/set", bytes.NewBuffer(reqJson))
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	testRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(setEntry)
	handler.ServeHTTP(testRecorder, req)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}
	
	var result interface{}
	errJson := json.NewDecoder(testRecorder.Body).Decode(&result)
	if errJson == nil {
		t.Fail()
		t.Logf(errJson.Error())
	}

	if result != nil {
		t.Fail()
		t.Logf(fmt.Sprint("result should be nil but found: ", result))
	}
}

func TestGetEntryRequest(t *testing.T) {
	reqJson, errReqJson := json.Marshal(testKind)
	if errReqJson != nil {
		t.Fail()
		t.Logf(errReqJson.Error())
		return
	}

	req, errReq := http.NewRequest("GET", "/get", bytes.NewBuffer(reqJson))
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	testRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(getEntry)
	handler.ServeHTTP(testRecorder, req)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}
	
	var result []byte
	errJson := json.NewDecoder(testRecorder.Body).Decode(&result)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
	}

	resultString := string(result)
	if resultString != testMessage {
		t.Fail()
		t.Logf(fmt.Sprint("result should be: ", testMessage, ", but found: ", result))
	}
}