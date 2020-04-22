package errors

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/responses"
)

var (
	BadBodyFail = "unable to decode request body"
	UnrecognizedQuery = "unrecognized query action requested"
	UnrecognizedMutation = "unrecognized mutation action requested"
	UnableToCreatePublicSession = "unable to create public session"
	InvalidSessionCredentials = "invalid session credentials provided"
	UnableToUpdateSession = "unable to update session"
	SessionProvidedIsNil = "session provided is nil"
	CredentialsProvidedAreNil = "credentials provided are nil"
)

var defaultFail = "unable to return session"

func DefaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	BadRequest(w, &responses.ErrorsPayload{
		Default: &errAsStr,
	})
}

func CustomErrorResponse(w http.ResponseWriter, err string) {
	BadRequest(w, &responses.ErrorsPayload{
		Default: &err,
	})
}

func BadRequest(w http.ResponseWriter, errors *responses.ErrorsPayload) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&responses.Body{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&responses.Body{
		Errors: &responses.ErrorsPayload{
			Default:	&defaultFail,
		},
	})
}
