package errors

import (
	"encoding/json"
	"net/http"
)

// Response -
type Response struct {
	Headers	 *string `json:"headers"`
	Body		 *string `json:"body"`
	Store		 *string `json:"store"`
	Session	 *string `json:"session"`
	Email    *string `json:"session"`
	Password *string `json:"password"`
	Default  *string `json:"default"`
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
func BadRequest(w http.ResponseWriter, errors *Response) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(errors)
	if err == nil {
		w.Write(body)
		return
	}

	failBody, _ := json.Marshal(&Response{
		Default: &defaultFail,
	})

	w.Write(failBody)
}
