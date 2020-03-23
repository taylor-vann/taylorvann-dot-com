package errors

import (
	"encoding/json"
	"net/http"
)

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
	Session	 *string `json:"session"`
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

var (
	// InvalidHeadersProvided -
	InvalidHeadersProvided = "invalid headers provided"
	// BadBodyFail -
	BadBodyFail = "unable to decode request body"
	// UnrecognizedQuery -
	UnrecognizedQuery = "unrecognized query action requested"
	// UnrecognizedMutation -
	UnrecognizedMutation = "unrecognized mutation action requested"
	// InvalidSessionCredentials -
	InvalidSessionCredentials = "invalid session credentials provided"
	// UnableToCreatePublicSession -
	UnableToCreatePublicSession = "unable to create public session"
	// UnableToCreateUser -
	UnableToCreateUser = "unable to create user"
	// UserAlreadyExists -
	UserAlreadyExists = "user already exists"
	// UserDoesNotExist -
	UserDoesNotExist = "user does not exist"
	// UnableToValidateUser -
	UnableToValidateUser = "unable to validate user"
)

var defaultFail = "unable to write custom error messages"

// BadRequest - mutate session whitelist
func BadRequest(w http.ResponseWriter, errors *Payload) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&ResponseBody{
			Errors: errors,
		})
		return
	}

	json.NewEncoder(w).Encode(&ResponseBody{
		Errors: &Payload{
			Default: &defaultFail,
		},
	})
}
