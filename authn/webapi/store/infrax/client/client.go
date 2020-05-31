package client

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

const (
	ApplicationJson = "application/json"
	ContentType = "Content-Type"
)

var Client, errClient = Init()

func Init() (*http.Client, error) {
	cookiejar, errCookiejar := cookiejar.New(
		&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		},
	)
	if errCookiejar != nil {
		return nil, errCookiejar
	}
	if cookiejar == nil {
		return nil, errors.New("nil jar provided")
	}

	client := &http.Client{
		Jar: cookiejar,
	}

	return client, nil
}

func Post(url string, domain *url.URL, payload *bytes.Buffer) (*http.Response, error) {
	request, errRequest := http.NewRequest("POST", url, payload)
	if errRequest != nil {
		return nil, errRequest
	}

	request.Header.Add("Content-Type", ApplicationJson)

	resp, errResp := Client.Do(request)
	Client.Jar.SetCookies(domain, resp.Cookies())

	return resp, errResp
}