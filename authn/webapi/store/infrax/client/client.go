package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"webapi/store/infrax/fetch/requests"
	"webapi/store/infrax/fetch/responses"

	"golang.org/x/net/publicsuffix"
)

const (
	ApplicationJson = "application/json"
	ContentType = "Content-Type"
	Domain = "https://briantaylorvann.com"
	SessionCookieHeader = "briantaylorvann.com_session"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
)

var (
	Environemnt = os.Getenv("STAGE")

	infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")

	guestSessionRequestBody = requests.Body{
		Action: "CREATE_GUEST_SESSION",
		Params: requests.GuestSessionParams {
			Environment: Environemnt,
		},
	}

	infraSessionRequestBody = requests.Body{
		Action: "CREATE_INFRA_OVERLORD_SESSION",
		Params: requests.InfraSession {
			Environment: Environemnt,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
	}
)

var ParsedDomain, errParsedDomain = url.Parse(Domain)
var Client, errClient = Init()
var Session, errSession = Setup()

func getRequestBodyBuffer(item interface{}) (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(item)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func Post(url string, payload *bytes.Buffer) (*http.Response, error) {
	if errClient != nil {
		return nil, errClient
	}
	if errParsedDomain != nil {
		return nil, errParsedDomain
	}

	request, errRequest := http.NewRequest("POST", url, payload)
	if errRequest != nil {
		return nil, errRequest
	}
	request.Header.Add("Content-Type", ApplicationJson)

	resp, errResp := Client.Do(request)
	if errResp != nil {
		return nil, errResp
	}
	
	Client.Jar.SetCookies(ParsedDomain, resp.Cookies())

	return resp, errResp
}

func GuestSession() (string, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		guestSessionRequestBody,
	)
	if errRequestBodyBuffer != nil {
		return "", errRequestBodyBuffer
	}

	resp, errResp := Post(
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errResp != nil {
		return "", errResp
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return "", errJson
	}
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}
	if responseBody.Session != nil {
		return responseBody.Session.Token, nil
	}

	return  "", errors.New("nil session returned")
}

func InfraSession() (string, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		infraSessionRequestBody,
	)
	if errRequestBodyBuffer != nil {
		return "", errRequestBodyBuffer
	}

	resp, errResp := Post(
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errResp != nil {
		return "", errResp
	}
	
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return "", errJson
	}
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}
	if responseBody.Session != nil {
		return responseBody.Session.Token, nil
	}

	return  "", errors.New("nil session returned")
}

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

func Setup() (*http.Cookie, error) {
	guestSession, errGuestSession := GuestSession()
	if errGuestSession != nil {
		
		return nil, errGuestSession
	}
	if guestSession == "" {
		
		return nil, errors.New("nil guest session returned")
	}

	infraSession, errInfraSession := InfraSession()
	if errInfraSession != nil {
		return nil, errInfraSession

	}
	if infraSession == "" {
		return nil, errors.New("nil infra session returned")
	}
	
	for _, cookie := range Client.Jar.Cookies(ParsedDomain) {
		if cookie.Name == SessionCookieHeader {
			return cookie, nil
		}
	}
	
	return nil, errors.New("did not find session cookie")
}