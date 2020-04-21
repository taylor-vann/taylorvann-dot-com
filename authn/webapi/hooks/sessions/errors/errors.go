package errors

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/responses"
)

type Credentials struct {
	Email			string		`json:"email"`
	Password	string		`json:"password"`
}

type RequestPayload struct {
	SessionToken	string			`json:"session_token"`
	Credentials		Credentials	`json:"credentials"`
}

type RequestBody struct {
	Action string          `json:"action"`
	Params RequestPayload `json:"params"`
}

type SessionResponsePayload struct {
	SessionToken	string	`json:"session_token"`
	CsrfToken			string	`json:"csrf_token"`
}

type ResponsePayload struct {
	Headers	*string `json:"headers"`
	Body		*string `json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

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

var defaultFail = "unable to write custom error messages"

func DefaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	BadRequest(w, &responses.ErrorsResponsePayload{
		Default: &errAsStr,
	})
}

func CustomErrorResponse(w http.ResponseWriter, err string) {
	BadRequest(w, &responses.ErrorsResponsePayload{
		Default: &err,
	})
}

func BadRequest(w http.ResponseWriter, errors *responses.ErrorsResponsePayload) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&responses.ResponseBody{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&responses.ResponseBody{
		Errors: &responses.ErrorsResponsePayload{
			Default:	&defaultFail,
		},
	})
}
