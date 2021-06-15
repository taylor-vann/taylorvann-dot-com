package muxrouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"webapi/details"
)

const (
	testHTTPStr        = "http://awesome.sauce.com/yoyoyo?=hello=world#framented-fragement"
	testURLStr         = "https://awesome.sauce.com/yoyoyo?=hello=world#framented-fragement"
	expectedTestURLStr = "https://awesome.sauce.com/yoyoyo"
	superawesome       = "https://superawesome.com"
	expectedAddress    = "https://127.0.0.1:5000"
	exampleDetailsPath = "/usr/local/config/details.init.example.json"
)

func TestRedactURL(t *testing.T) {
	testURL, errTestURL := url.Parse(testURLStr)
	if errTestURL != nil {
		t.Fail()
		t.Logf(errTestURL.Error())
		return
	}

	redactedURL, errRedactedURL := RedactURL(testURL, nil)
	if errRedactedURL != nil {
		t.Fail()
		t.Logf(errRedactedURL.Error())
	}

	redactedURLStr := redactedURL.String()
	if redactedURLStr != expectedTestURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("example detail cert path: ", expectedTestURLStr, "\nfound:", redactedURLStr))
	}
}

func TestRedactURLFromString(t *testing.T) {
	redactedURL, errRedactedURL := RedactURLFromString(testURLStr, nil)
	if errRedactedURL != nil {
		t.Fail()
		t.Logf(errRedactedURL.Error())
		return
	}

	redactedURLStr := redactedURL.String()
	if redactedURLStr != expectedTestURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("example detail cert path: ", expectedTestURLStr, "\nfound:", redactedURLStr))
	}
}

func TestRedirectToHTTPS(t *testing.T) {
	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	testRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(redirectToHTTPS)
	handler.ServeHTTP(testRecorder, request)

	locations := testRecorder.HeaderMap["Location"]
	if locations[0] != testURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
	}
}

func TestCreateRedirectToHttpsMux(t *testing.T) {
	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	testRecorder := httptest.NewRecorder()

	mux := CreateRedirectToHttpsMux()
	mux.ServeHTTP(testRecorder, request)

	locations := testRecorder.HeaderMap["Location"]
	if locations[0] != testURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
	}
}

func TestCreateProxyMux(t *testing.T) {
	exampleDetails, errExampleDetails := details.ReadDetailsFromFile(exampleDetailsPath)
	if errExampleDetails != nil {
		t.Fail()
		t.Logf(errExampleDetails.Error())
	}

	proxyMux, errProxyMux := CreateProxyMux(&exampleDetails.Routes)
	if errProxyMux != nil {
		t.Fail()
		t.Logf(errProxyMux.Error())
	}

	if proxyMux == nil {
		t.Fail()
		t.Logf("proxyMux was not created")
	}
}
