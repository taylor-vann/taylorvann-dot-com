package client

import (
	"log"
	"testing"
	"net/url"
)

func TestInit(t *testing.T) {
	Init()

	if CookieJar == nil {
		t.Error("cookie jar is nil")
	}
	if Client == nil {
		t.Error("cookie jar is nil")
	}
}

func TestInitHasCookiePresent(t *testing.T) {
	urlStr, errUrlStr := url.Parse(CookieDomain)
	if errUrlStr != nil {
		t.Error(errUrlStr)
	}

	for _, cookie := range CookieJar.Cookies(urlStr) {
		log.Println(cookie.Name, cookie.Value)
	}
	if CookieJar == nil {
		t.Error("cookie jar is nil")
	}
	if Client == nil {
		t.Error("cookie jar is nil")
	}
}