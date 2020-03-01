package mutations

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"webapi/hooks/constants"
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
	"webapi/store"
)

// CreateUserRequestBodyParams -
type CreateUserRequestBodyParams struct {
	Action string                 `json:"action"`
	Params store.CreateUserParams `json:"params"`
}

func defaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	errors.BadRequest(w, &errors.Response{
		Default: &errAsStr,
	})
}

// CreatePublicSession - mutate session whitelist
func CreatePublicSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("creating public session")
	if !validateGuestHeaders(r) {
		errors.BadRequest(w, &errors.Response{
			Session: &InvalidHeadersProvided,
		})
		return
	}

	// different name
	var body CreateUserRequestBodyParams
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		defaultErrorResponse(w, err)
		return
	}

	userSessionToken, errUserSessionToken := sessions.ComposeCreatePublicSessionParams(
		&sessions.CreatePublicJWTParams{
			Email:    body.Params.Email,
			Password: body.Params.Password,
		},
	)
	// create guest session
	if errUserSessionToken != nil {
		errorAsStr := errUserSessionToken.Error()
		errors.BadRequest(w, &errors.Response{
			Session: &errors.InvalidSessionCredentials,
			Default: &errorAsStr,
		})
		return
	}

	userSession, errUserSession := sessions.Create(
		userSessionToken,
	)

	if errUserSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(userSession.CsrfToken)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set(constants.SessionTokenHeader, userSession.SessionToken)
		w.Header().Set(constants.CsrfTokenHeader, csrfAsBase64)
		w.WriteHeader(http.StatusOK)
		return
	}

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &errors.Response{
		Session: &errors.UnableToCreatePublicSession,
		Default: &errorAsStr,
	})
}
