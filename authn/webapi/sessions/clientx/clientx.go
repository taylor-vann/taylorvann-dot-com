package clientx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"webapi/sessions/clientx/infrafetchx"
	"webapi/sessions/clientx/infrafetchx/requests"
	"webapi/sessions/clientx/infrafetchx/responses"
	"webapi/sessions/clientx/sessionx"
)

var InfraSessionCookie *http.Cookie
var errInfraSessionCookie error

var httpClient = http.Client{}

func Setup() (*http.Cookie, error) {
	InfraSessionCookie, errInfraSessionCookie = sessionx.Setup()
	return InfraSessionCookie, errInfraSessionCookie
}

func ValidateSession(p requests.ValidateSession) (string, error) {
	body := requests.Body{
		Action: "VALIDATE_SESSION",
		Params: p,
	}

	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(body)
	if errJsonBuffer != nil {
		return "", errJsonBuffer
	}

	resp, errResp := Do(
		infrafetchx.SessionsQueryAddress,
		sessionBuffer,
	)
	if errResp != nil {
		return "", errResp
	}

	var respBody responses.SessionBody
	errDecode := json.NewDecoder(resp.Body).Decode(&respBody)
	if errDecode != nil {
		return "", errDecode
	}

	session := respBody.Session	
	if session != nil && session.Token != "" {
		return session.Token, nil
	}
	
	return "", errors.New("could not verify session")
}

func Do(address string, payload *bytes.Buffer) (*http.Response, error) {
	if InfraSessionCookie == nil || errInfraSessionCookie != nil {
		return nil, errors.New("no internal session provided")
	}

	req, errReq := http.NewRequest("POST", address, payload)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(InfraSessionCookie)

	return httpClient.Do(req)
}