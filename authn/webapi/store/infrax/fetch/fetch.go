package fetch

import (
	"net/http"

	"bytes"
	"encoding/json"
	"errors"
	"os"

	"log"

	"webapi/store/infrax/requests"
	"webapi/store/infrax/responses"
)

const (
	ApplicationJson = "application/json"
	AuthnSessionQueryAddress = "https://authn.briantaylorvann.com/q/sessions/"
	AuthnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
	GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 
)

const CookieDomain = "briantaylorvann.com"
const InternalSessionCookieHeader = "briantaylorvann.com_internal_session"

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

func GetInternalGuestSessionCookie(sessionToken string) *http.Cookie {
	return &http.Cookie{
		Name:			InternalSessionCookieHeader,
		Value:		sessionToken,
		MaxAge:		GuestSessionExpirationInSeconds,
		Domain:   CookieDomain,
		Path:     "/",
		Secure:		true,
		HttpOnly:	true,
		SameSite:	3,
	}
}

func GuestSession() (string, error) {
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
		log.Println(errResp.Error())
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.Body
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		log.Println(string(resp.StatusCode))
		log.Println(errJson)
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