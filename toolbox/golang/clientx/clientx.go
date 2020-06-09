package clientx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/weblog/toolbox/golang/clientx/fetch"
	"github.com/weblog/toolbox/golang/clientx/fetch/requests"
	"github.com/weblog/toolbox/golang/clientx/fetch/responses"
	"github.com/weblog/toolbox/golang/clientx/sessionx"
)

var httpClient = http.Client{}

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
		fetch.SessionsQueryAddress,
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