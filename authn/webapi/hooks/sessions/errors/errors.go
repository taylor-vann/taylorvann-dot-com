package errors

import (
	"encoding/json"
	"net/http"
)

type Credentials struct {
	Email			*string		`json:"email"`
	Password	*string		`json:"password"`
}

type RequestPayload struct {
	SessionToken	*string				`json:"session_token"`
	CsrfToken			*string				`json:"csrf_token"`
	Credentials		*Credentials	`json:"credentials"`
}

type RequestBody struct {
	Action string          `json:"action"`
	Params *RequestPayload `json:"params"`
}

type SessionResponsePayload struct {
	SessionToken	*string	`json:"session_token"`
	CsrfToken			*string	`json:"csrf_token"`
}

type ResponsePayload struct {
	Headers	*string	`json:"headers"`
	Body		*string	`json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

type ResponseBody struct {
	Session *SessionResponsePayload	`json:"session"`
	Errors  *ResponsePayload				`json:"errors"`
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
	BadRequest(w, &ResponsePayload{
		Default: &errAsStr,
	})
}

func CustomErrorResponse(w http.ResponseWriter, err string) {
	BadRequest(w, &ResponsePayload{
		Default: &err,
	})
}

func BadRequest(w http.ResponseWriter, errors *ResponsePayload) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&ResponseBody{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&ResponseBody{
		Errors: &ResponsePayload{
			Default:	&defaultFail,
		},
	})
}
