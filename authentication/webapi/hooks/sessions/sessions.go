package sessions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/mutations"
)

// ReadSessionAction -
type ReadSessionAction struct {
	SessionSignature string `json:"session_signature"`
}

// RemoveSessionAction -
type RemoveSessionAction = ReadSessionAction

// RequestBodyParams -
type RequestBodyParams struct {
	Action string
	Params interface{}
}

// Actions
const (
	// CreateGuestSession -
	CreateGuestSession = "CREATE_GUEST_SESSION"
	// CreatePublicSession -
	CreatePublicSession = "CREATE_PUBLIC_SESSION"
	// CreatePublicPasswordResetSession -
	CreatePublicPasswordResetSession = "CREATE_PUBLIC_PASSWORD_RESET_SESSION"
	// ReadSession -
	ReadSession = "READ_SESSION"
	// ValidateSession -
	ValidateSession = "VALIDATE_SESSION"
	// RemoveSession -
	RemoveSession = "REMOVE_SESSION"
)

// Query - read information from session whitelist
func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Response{
			Body: &errors.BadBodyFail,
		})
		return
	}

	var body RequestBodyParams
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errAsStr := err.Error()
		errors.BadRequest(w, &errors.Response{
			Session: &errors.BadBodyFail,
			Default: &errAsStr,
		})
	}

	switch body.Action {
	case ReadSession:
	default:
		errors.BadRequest(w, &errors.Response{
			Session: &errors.UnrecognizedQuery,
		})
	}
}

// Mutation - mutate session whitelist
func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Response{
			Body: &errors.BadBodyFail,
		})
		return
	}

	var body RequestBodyParams
	if r.Body != nil {
		b, _ := ioutil.ReadAll(r.Body)

		err := json.NewDecoder(
			ioutil.NopCloser(bytes.NewReader(b)),
		).Decode(&body)

		if err != nil {
			errors.BadRequest(w, &errors.Response{
				Default: &errors.BadBodyFail,
			})
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	switch body.Action {
	case CreateGuestSession:
		mutations.CreateGuestSession(w, r)
	case CreatePublicSession:
		fmt.Println("create public session action")
		mutations.CreatePublicSession(w, r)
	case CreatePublicPasswordResetSession:
		mutations.CreatePublicPasswordResetSession(w, r)
	// case ValidateSession:
	case RemoveSession:
		mutations.RemoveSession(w, r)
	default:
		errors.BadRequest(w, &errors.Response{
			Session: &errors.UnrecognizedMutation,
		})
	}
}
