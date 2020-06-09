package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/requests"
	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/responses"
)

const (
	UsersStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	RolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	SessionsQueryAddress = "https://authn.briantaylorvann.com/q/sessions/"
)

var (
	Environemnt = os.Getenv("STAGE")

	client = http.Client{}
)

func getRequestBodyBuffer(item interface{}) (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(item)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func ValidateGuestSession(p requests.ValidateGuestSession, sessionCookie *http.Cookie) (string, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_GUEST_SESSION",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return "", errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
		SessionsQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return "", errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
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

func ValidateGuestUser(p requests.ValidateGuestUser, sessionCookie *http.Cookie) (*responses.User, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_GUEST_USER",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
		UsersStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errors.New("errors were returned in fetch")
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return  nil, errors.New("nil session returned")
}

func ValidateInfraRole(p requests.ValidateInfraRole, sessionCookie *http.Cookie) (*responses.Role, error) {
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_INFRA_OVERLORD_ROLE",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.RolesBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errors.New("errors were returned in fetch")
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}
	
	return nil, errors.New("unable to validate infra role")
}
