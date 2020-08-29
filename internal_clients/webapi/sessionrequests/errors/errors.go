package errors

import (
	"encoding/json"
	"net/http"

	"webapi/sessionrequests/responses"
)

var (
	BadRequestFail  		 			= "unable to decode request body"
	UnrecognizedParams				= "unrecognized params given"
	FailedToValidateGuestUser	= "failed to validate guest user"
)

var defaultFail = "unable to return Roles"

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
