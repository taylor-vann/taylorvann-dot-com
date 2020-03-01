package mutations

import (
	"encoding/base64"
	"net/http"
	"webapi/hooks/constants"
	"webapi/interfaces/jwtx"
	"webapi/sessions"
	sessionsc "webapi/sessions/constants"
)

// BadSessionRequestErrResponse -
type BadSessionRequestErrResponse struct {
	Session string `json:"session"`
}

// RemoveSessionRequestParams -
type RemoveSessionRequestParams struct {
	Signature string `json:"session_signature"`
}

// RemoveSessionRequest -
type RemoveSessionRequest struct {
	Action string `json:"action"`
}

// ResponseErrors -
type ResponseErrors struct {
	Session *string `json:"session"`
	Default *string `json:"default"`
}

// ResponseBody -
type ResponseBody struct {
	Errors *ResponseErrors `json:"errors"`
}

// CreateGuestSessionErrorMessage -
var CreateGuestSessionErrorMessage = "error creating guest session"

// InvalidHeadersProvided -
var InvalidHeadersProvided = "invalid headers provided"

func validateGuestSessionToken(token *string) bool {
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
