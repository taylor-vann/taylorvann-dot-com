package errors

import (
	"encoding/json"
	"net/http"
)

// Response -
type Response struct {
	Headers	*string	`json:"headers"`
	Body		*string	`json:"body"`
	Session *string `json:"session"`
	Default *string `json:"default"`
}

var (
	// BadBodyFail - 
	BadBodyFail = "unable to decode request body"
	// UnrecognizedQuery -
	UnrecognizedQuery = "unrecognized query action requested"
	// UnrecognizedMutation -
	UnrecognizedMutation = "unrecognized mutation action requested"
	// UnableToCreatePublicSession -
	UnableToCreatePublicSession = "unable to create public session"
	// InvalidSessionCredentials -
	InvalidSessionCredentials = "invalid session credentials provided"
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
		Default:	&defaultFail,
	})

	w.Write(failBody)
}
