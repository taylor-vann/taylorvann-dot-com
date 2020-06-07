package validatesessionx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"webapi/store/validatesessionx/requests"
	"webapi/store/validatesessionx/responses"

)

const (
	ApplicationJson = "application/json"
	AuthnSessionQueryAddress = "https://authn.briantaylorvann.com/q/sessions/"
	AuthnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
)

var (
	Environment = os.Getenv("STAGE")
	RequestGuestSessionBody = requests.Body{
		Action: "CREATE_GUEST_SESSION",
		Params: requests.GuestSessionParams {
			Environment: "LOCAL",
		},
	}
)

func getGuestSessionRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestGuestSessionBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func getValidateSessionRequestBodyBuffer(sessionToken string) (*bytes.Buffer, error) {
	var RequestValidateSessionBody = requests.Body{
		Action: "VAIDATE_GUEST_SESSION",
		Params: requests.ValidateParams {
			Environment: "LOCAL",
			Token: sessionToken,
		},
	}
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestValidateSessionBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func FetchGuestSession() (string, error) {
	var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getGuestSessionRequestBodyBuffer()
	if errGuestSessionRequestBodyBuffer != nil {
		return "", errGuestSessionRequestBodyBuffer
	}

	resp, errResp := http.Post(
		AuthnSessionMutationAddress,
		ApplicationJson,
		guestSessionRequestBodyBuffer,
	)
	if errResp != nil {
		return "", errResp
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.Body
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

func ValidateGuestSession(sessionToken string) (bool, error) {
	validateGuestBuffer, errValidateSessionBuffer := getValidateSessionRequestBodyBuffer(sessionToken)
	if errValidateSessionBuffer != nil {
		return false, errValidateSessionBuffer
	}
	resp, errResp := http.Post(
		AuthnSessionQueryAddress,
		ApplicationJson,
		validateGuestBuffer,
	)
	if errResp != nil {
		return false, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return false, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.Body
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return false, errJson
	}
	if responseBody.Errors != nil {
			return false, errors.New("errors were returned in fetch")
	}

	if responseBody.Session != nil {
		return true, nil
	}

	return false,  errors.New("nil session returned")
}