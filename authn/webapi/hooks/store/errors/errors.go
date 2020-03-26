package errors

import (
	"encoding/json"
	"net/http"
)

// problematic because user / roles are different

// SessionParams -
type SessionParams struct {
	SessionToken	*string	`json:"session_token"`
	CsrfToken			*string	`json:"csrf_token"`
}

// ResponsePayload -
type ResponsePayload struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

// Payload -
type Payload struct {
	Headers  *string `json:"headers"`
	Body     *string `json:"body"`
	Store    *string `json:"store"`
	Email    *string `json:"session"`
	Password *string `json:"password"`
	Default  *string `json:"default"`
}

// ResponseBody -
type ResponseBody struct {
	User		*ResponsePayload `json:"user"`
	Errors	*Payload       	 `json:"errors"`
}

const (
	BadRequest = "unable to decode request body"
	UnrecognizedQuery = "unrecognized query action requested"
	UnrecognizedMutation = "unrecognized mutation action requested"
	InvalidSessionCredentials = "invalid session credentials provided"
	UnableToCreatePublicSession = "unable to create public session"
	UnableToCreateUser = "unable to create user"
	UserAlreadyExists = "user already exists"
	UserDoesNotExist = "user does not exist"
	UnableToValidateUser = "unable to validate user"
	DefaultFail = "unable to write custom error messages"
)

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
			Default:	&DefaultFail,
		},
	})
}

