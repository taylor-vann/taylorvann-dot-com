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
	AuthnUserStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	AuthnRolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	AuthnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
	GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3 
)

const CookieDomain = "briantaylorvann.com"
const SessionCookieHeader = "briantaylorvann.com_session"

var (
	Environment = os.Getenv("STAGE")
	InfraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	InfraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")

	RequestValidateUserBody = requests.Body{
		Action: "VALIDATE_USER",
		Params: requests.ValidateUserParams {
			Environment: "LOCAL",
			Email: InfraOverlordEmail,
			Password: InfraOverlordPassword,
		},
	}
)

func getValidateUserRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestValidateUserBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

// our needs

// validate guest user (requires guest session)
// validate guest user internal role (requires guest session)
// get internal overlord session (requires guest session)

// func CreateValidateUserRequest()  {
// 	log.Println("caalled to fail :)"
// 	var validateUserRequestBodyBuffer, errValidateRequestBodyBuffer = GetValidateUserRequestBodyBuffer()
// 	if errValidateRequestBodyBuffer != nil {
// 		return "", errValidateRequestBodyBuffer
// 	}

// 	resp, errResp := http.NewRequest(
// 		"POST"
// 		AuthnSessionMutationAddress,
// 		validateUserRequestBodyBuffer,
// 	)
// }