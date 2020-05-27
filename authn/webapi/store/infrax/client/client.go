package client

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"webapi/store/infrax/fetch"
)

const CookieDomain = "https://briantaylorvann.com/"
var CookieHeader = "briantaylorvann.com_internal_session"

var CookieJar *cookiejar.Jar
var Client *http.Client

// fetch guest session
func Init() (*http.Client, error) {
	cookiejar, errCookiejar := cookiejar.New(nil)
	if errCookiejar != nil {
		return nil, errCookiejar
	}
	CookieJar = cookiejar
	if cookiejar == nil {
		return nil, errors.New("nil jar provided")
	}
	Client = &http.Client{
		Jar: cookiejar,
	}

	session, errSession := fetch.GuestSession()
	if errSession != nil {
		return nil, errSession
	}

	var cookies []*http.Cookie
	cookies = append(cookies, fetch.GetInternalGuestSessionCookie(session))
	urlStr, errUrlStr := url.Parse(CookieDomain)
	if errUrlStr != nil {
		return nil, errUrlStr
	}
	cookiejar.SetCookies(urlStr, cookies)

	return Client, nil
}

// store guest session

// validate user with session

// change to validate internal role

func Post(urlStr string, payload *bytes.Buffer) (*http.Response, error) {
	request, errRequest := http.NewRequest("POST", urlStr, payload)
	if errRequest != nil {
		return nil, errRequest
	}

	request.Header.Add("Content-Type", "application/json")

	return Client.Do(request)
}