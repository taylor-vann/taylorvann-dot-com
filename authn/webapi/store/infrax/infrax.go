import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"log"

	"webapi/store/infrax/client"
	"webapi/store/infrax/requests"
	"webapi/store/infrax/responses"

	"golang.org/x/net/publicsuffix"
)

const (
	authnUserStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	authnRolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	authnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	domain = "https://briantaylorvann.com"
)

var (
	Environemnt = os.Getenv("STAGE")
	Client = client.Client
)

var (
	infraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	infraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")

	requestGuestSessionBody = requests.Body{
		Action: "CREATE_GUEST_SESSION",
		Params: requests.GuestSessionParams {
			Environment: Environemnt,
		},
	}

	requestValidateUserBody = requests.Body{
		Action: "VALIDATE_GUEST_USER",
		Params: requests.ValidateUserParams {
			Environment: Environemnt,
			Email: infraOverlordEmail,
			Password: infraOverlordPassword,
		},
	}
)

var parsedDomain, errParsedDomain = url.Parse(domain)

func getGuestSessionRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(requestGuestSessionBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func getValidateUserRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(requestValidateUserBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func fetchGuestSession() (string, error) {
	var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getGuestSessionRequestBodyBuffer()
	if errGuestSessionRequestBodyBuffer != nil {
		return "", errGuestSessionRequestBodyBuffer
	}

	resp, errResp := client.Post(
		authnSessionMutationAddress,
		parsedDomain,
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

// go for guest session to prove it first
func fetchValidateUser() (string, error) {
	var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getValidateUserRequestBodyBuffer()
	if errGuestSessionRequestBodyBuffer != nil {
		return "", errGuestSessionRequestBodyBuffer
	}

	resp, errResp := client.Post(
		authnUserStoreQueryAddress,
		parsedDomain,
		guestSessionRequestBodyBuffer,
	)
	if errResp != nil {
		return "", errResp
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return "", errJson
	}
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}
	if responseBody.Users != nil {
		return "found a user", nil
	}

	return  "", errors.New("nil session returned")
}

// fetch validate internal role

// create infra admin session

// setup
// create infra admin session