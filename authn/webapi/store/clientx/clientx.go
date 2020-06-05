package clientx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"webapi/store/clientx/fetch/requests"
	"webapi/store/clientx/sessionx"
)

var httpClient = http.Client{}

func ValidateSession(p requests.ValidateSession) (*http.Response, error) {
	params := requests.ValidateSession{
		Environment: p.Environment,
		Token: p.Token,
	}

	body := requests.Body{
		Action: "VALIDATE_SESSION",
		Params: params,
	}

	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(body)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return Do(
		"https://authn.briantaylorvann.com/q/sessions",
		sessionBuffer,
	)
}

func Do(address string, payload *bytes.Buffer) (*http.Response, error) {
	if sessionx.Session == nil {
		return nil, errors.New("no internal session provided")
	}

	req, errReq := http.NewRequest(
		"POST",
		address,
		payload,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionx.Session)

	return httpClient.Do(req)
}