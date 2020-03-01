package mutations

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"errors"
	"webapi/hooks/constants"
	storee "webapi/hooks/store/errors"
	"webapi/interfaces/jwtx"
	"webapi/sessions"
	sessionsc "webapi/sessions/constants"

	"webapi/controllers/users"
)

// CreateUserParams -
type CreateUserParams = users.CreateParams

// CreateUserRequestBody -
type CreateUserRequestBody struct {
	Action string           `json:"action"`
	Params CreateUserParams `json:"params`
}

func defaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	storee.BadRequest(w, &storee.Response{
		Default: &errAsStr,
	})
}

func validateGuestSessionToken(token *string) bool {
	fmt.Println("validate geust session token")
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(token)
	if errTokenDetails != nil {
		return false
	}
	if tokenDetails.Payload.Iss != sessionsc.TaylorVannDotCom {
		return false
	}
	if tokenDetails.Payload.Sub != sessionsc.Guest {
		return false
	}
	if tokenDetails.Payload.Aud != sessionsc.Public {
		return false
	}
	return true
}

// validateGuestHeaders - valid guest session and csrf are required
func validateGuestHeaders(r *http.Request) bool {
	fmt.Println("validating guest headers")
	sessionToken := r.Header.Get(constants.SessionTokenHeader)
	if sessionToken == "" {
		return false
	}
	csrfTokenBase64 := r.Header.Get(constants.CsrfTokenHeader)
	if csrfTokenBase64 == "" {
		return false
	}

	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(csrfTokenBase64)
	if errCsrfToken != nil {
		return false
	}

	result, errEntry := sessions.Check(&sessions.CheckParams{
		SessionToken: &sessionToken,
		CsrfToken:    &csrfToken,
	})
	if errEntry != nil {
		return false
	}

	if result != nil {
		// check for guest session
		return validateGuestSessionToken(&sessionToken)
	}

	return false
}


// validateGuestHeaders - valid guest session and csrf are required
func validatePublicHeaders(r *http.Request) (*sessions.Session, error) {
	fmt.Println("validating guest headers")
	sessionToken := r.Header.Get(constants.SessionTokenHeader)
	if sessionToken == "" {
		return nil, errors.New("")
	}
	csrfTokenBase64 := r.Header.Get(constants.CsrfTokenHeader)
	if csrfTokenBase64 == "" {
		return nil, errors.New("")
	}

	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(csrfTokenBase64)
	if errCsrfToken != nil {
		return nil, errors.New("")
	}

	return sessions.Update(&sessions.UpdateParams{
		SessionToken: &sessionToken,
		CsrfToken:    &csrfToken,
	})
}