package errors

import (
	"encoding/json"
	"net/http"

	"webapi/infraclientx/verifyx/responses"
)

const (
	ContentType = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	BadRequestFail = "unable to decode request body"
	InvalidGuestSession	= "invalid guest session"
	InvalidInfraSession	= "invalid infra session provided"

	defaultFail = "unable to return Roles"
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
	w.Header().Set(ContentType, ApplicationJson)

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
