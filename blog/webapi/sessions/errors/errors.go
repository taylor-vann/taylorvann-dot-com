package errors

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/responses"
)

var (
	BadRequestFail     = "unable to decode request body"
	UnrecognizedParams = "unrecognized params given"
)

var defaultFail = "unable to return a session"

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