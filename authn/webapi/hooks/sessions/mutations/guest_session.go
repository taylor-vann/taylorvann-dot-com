package mutations

import (
	"encoding/base64"
	// "fmt"
	"net/http"
	"webapi/hooks/constants"
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

// validateGuestSessionHeaders - make sure no csrf or session token exists
func validateGuestSessionHeaders(r *http.Request) bool {
	sessionToken := r.Header.Get(constants.SessionTokenHeader)
	if sessionToken != "" {
		return false
	}
	csrfToken := r.Header.Get(constants.CsrfTokenHeader)
	if csrfToken != "" {
		return false
	}

	return true
}

// CreateGuestSession - mutate session whitelist
func CreateGuestSession(w http.ResponseWriter, r *http.Request) {
	if !validateGuestSessionHeaders(r) {
		errors.BadRequest(w, &errors.Response{
			Session: &InvalidHeadersProvided,
		})
		return
	}

	session, errSession := sessions.Create(
		sessions.ComposeCreateGuestSessionParams(),
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
