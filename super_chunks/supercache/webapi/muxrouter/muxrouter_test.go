package muxrouter

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"

// 	"webapi/details"
// )

// func TestRedirectToHTTPS(t *testing.T) {
// 	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
// 	if errRequest != nil {
// 		t.Fail()
// 		t.Logf(errRequest.Error())
// 	}

// 	testRecorder := httptest.NewRecorder()

// 	handler := http.HandlerFunc(redirectToHTTPS)
// 	handler.ServeHTTP(testRecorder, request)

// 	locations := testRecorder.HeaderMap["Location"]
// 	if locations[0] != testURLStr {
// 		t.Fail()
// 		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
// 	}
// }

// func TestCreateRedirectToHttpsMux(t *testing.T) {
// 	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
// 	if errRequest != nil {
// 		t.Fail()
// 		t.Logf(errRequest.Error())
// 	}

// 	testRecorder := httptest.NewRecorder()

// 	mux := CreateRedirectToHttpsMux()
// 	mux.ServeHTTP(testRecorder, request)

// 	locations := testRecorder.HeaderMap["Location"]
// 	if locations[0] != testURLStr {
// 		t.Fail()
// 		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
// 	}
// }

// func TestCreateProxyMux(t *testing.T) {
// 	exampleDetails, errExampleDetails := details.ReadDetailsFromFile(exampleDetailsPath)
// 	if errExampleDetails != nil {
// 		t.Fail()
// 		t.Logf(errExampleDetails.Error())
// 	}

// 	proxyMux, errProxyMux := CreateProxyMux(&exampleDetails.Routes)
// 	if errProxyMux != nil {
// 		t.Fail()
// 		t.Logf(errProxyMux.Error())
// 	}

// 	if proxyMux == nil {
// 		t.Fail()
// 		t.Logf("proxyMux was not created")
// 	}
// }
