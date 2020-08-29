package errors

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/responses"
)

var (
	BadRequestFail       = "unable to decode request body"
	NilRequestBodyFail   = "request body is nil"
	UnrecognizedQuery    = "unrecognized query action requested"
	UnrecognizedMutation = "unrecognized mutation action requested"

	UnableToValidateSession    = "unable to create validate session"
	InvalidSessionCredentials  = "invalid session credentials provided"
	InvalidSessionProvided     = "invalid session provided"
	InvalidDefaultUserProvided = "invalid default user provided"

	defaultFail = "unable to return session"
)

func DefaultResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	BadRequest(w, &responses.Errors{
		Default: &errAsStr,
	})
}

func CustomResponse(w http.ResponseWriter, err string) {
	BadRequest(w, &responses.Errors{
		Default: &err,
	})
}

func BadRequest(w http.ResponseWriter, errors *responses.Errors) {
	w.WriteHeader(http.StatusBadRequest)

	if errors != nil {
		json.NewEncoder(w).Encode(&responses.Body{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&responses.Body{
		Errors: &responses.Errors{
			Default: &defaultFail,
		},
	})
}
