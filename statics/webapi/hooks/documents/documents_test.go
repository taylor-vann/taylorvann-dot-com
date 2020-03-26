package documents

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepageNoTokens(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp)
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Homepage)
	handler.ServeHTTP(httpTest, resp)

	if httpTest.Code != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}

	t.Error("fail because we're new")
}

func TestLoginNoTokens(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/",
		nil,
	)
	if errResp != nil {
		t.Error(errResp)
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(httpTest, resp)

	if httpTest.Code != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}

	t.Error("fail because we're new")
}
