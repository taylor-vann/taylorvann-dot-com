package infrax

import (
	"errors"
	"net/http"

	"webapi/store/infrax/client"
)

const httpClient = http.Client{}

func Do(address string, payload *bytes.Buffer) (*http.Response, error) {
	if client.Session == nil {
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
	req.AddCookie(client.Session)

	return client.Do(req)
}