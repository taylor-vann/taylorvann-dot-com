package mutations

import (
	"encoding/base64"
	// "fmt"
	"net/http"
	"webapi/hooks/constants"
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

// CreatePublicPasswordResetSession - mutate session whitelist
func CreatePublicPasswordResetSession(w http.ResponseWriter, r *http.Request) {
	if !validateGuestHeaders(r) {
		errors.BadRequest(w, &errors.Response{
			Session: &InvalidHeadersProvided,
		})
		return
	}

	session, errSession := sessions.Create(
		sessions.ComposeCreateResetPasswordSessionParams(),
	)

	if errSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set(constants.SessionTokenHeader, session.SessionToken)
		w.Header().Set(constants.CsrfTokenHeader, csrfAsBase64)
		w.WriteHeader(http.StatusOK)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &errors.Response{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
