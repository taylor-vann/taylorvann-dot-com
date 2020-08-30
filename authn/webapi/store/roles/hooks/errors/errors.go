package errors

import (
	"encoding/json"
	"net/http"

	"webapi/store/roles/hooks/responses"
)

var (
	BadRequestFail       = "unable to decode request body"
	UnrecognizedQuery    = "unrecognized query action requested"
	UnrecognizedMutation = "unrecognized mutation action requested"
	InvalidGuestSession  = "invalid guest session"
	NilInfraCredentials  = "nil infra session credentials provided"
	InvalidInfraSession  = "invalid infra session provided"

	FailedToCreateRole       = "failed to create role"
	FailedToReadRole         = "failed to read role"
	FailedToIndexRoles       = "failed to index roles"
	FailedToSearchRoles      = "failed to search roles"
	FailedToUpdateRole       = "failed to Update role"
	FailedToUpdateAccessRole = "failed to Update Access role"
	FailedToDeleteRole       = "failed to delete role"
	FailedToUndeleteRole     = "failed to undelete role"

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
