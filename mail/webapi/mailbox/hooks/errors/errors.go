// brian taylor vann
// taylorvann dot com

package errors

import (
	"encoding/json"
	"net/http"

	"webapi/mailbox/hooks/responses"
)

var (
	NilBodyGiven			 = "nil body given"
	BadRequestFail 	   = "unable to decode request body"
	UnrecognizedAction = "unrecognized action requested"
	UnrecognizedParams = "unrecognized parameters"
)

var defaultFail = "unable to send email"

func DefaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	BadRequest(w, &responses.Errors{
		Default: &errAsStr,
	})
}

func BadRequest(w http.ResponseWriter, errors *responses.Errors) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&responses.Body{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&responses.Body{
		Errors: &responses.Errors{
			Default:	&defaultFail,
		},
	})
}
