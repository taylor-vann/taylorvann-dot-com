package headersessionx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"webapi/headersessionx/cookies"
)

// requests
type GuestSessionParams struct {
	Environment string
}
type GuestSessionRequestBody struct {
	Action string
	Params GuestSessionParams
}

// responses
type Session struct {
	Token string
}

type Errors struct {
	Headers			*string `json:"headers"`
	RequestBody	*string `json:"request_body"`
	Session			*string `json:"session"`
	Default			*string `json:"default"`
}

type Body struct {
	Session *Session	`json:"session"`
	Errors  *Errors		`json:"errors"`
}

const (
	ApplicationJson = "application/json"
	UserSessionCookieHeader = "briantaylorvann.com_session"
	AuthnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
)

var Environment = os.Getenv("STAGE")

var requestGuestSessionBody = GuestSessionRequestBody{
	Action: "CREATE_GUEST_SESSION",
	Params: GuestSessionParams {
		Environment: "LOCAL",
	},
}

var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getGuestSessionRequestBodyBuffer()

func getGuestSessionRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(requestGuestSessionBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func FetchGuestSession() (string, error) {
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

	var responseBody Body
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return "", errJson
	}
	
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}

	if responseBody.Session == nil {
		return "", errors.New("nil session returned")
	}

	return responseBody.Session.Token, nil
}

func AttachGuestSession(w http.ResponseWriter) error {
	resp, errResp := FetchGuestSession()
	
	if errResp != nil {
		return errResp
	}

	cookies.AttachGuestSession(w, resp)
	
	return nil
}