package errors

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/responses"
)

var (
	BadRequestFail  		 = "unable to decode request body"
	UnrecognizedQuery 	 = "unrecognized query action requested"
	UnrecognizedMutation = "unrecognized mutation action requested"
	FailedToCreateRole 	 = "failed to create role"
	FailedToDeleteRole	 = "failed to update role"
	FailedToDeleteRole 	 = "failed to delete role"
	FailedToReviveRole 	 = "failed to revive role"
	FailedToIndexRoles 	 = "failed to index roles"
	FailedToSearchRoles  = "failed to search roles"
)

var defaultFail = "unable to return Roles"

func DefaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	BadRequest(w, &responses.ErrorsPayload{
		Default: &errAsStr,
	})
}

func CustomErrorResponse(w http.ResponseWriter, err string) {
	BadRequest(w, &responses.ErrorsPayload{
		Default: &err,
	})
}

func BadRequest(w http.ResponseWriter, errors *responses.ErrorsPayload) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	if errors != nil {
		json.NewEncoder(w).Encode(&responses.Body{Errors: errors})
		return
	}

	json.NewEncoder(w).Encode(&responses.Body{
		Errors: &responses.ErrorsPayload{
			Default:	&defaultFail,
		},
	})
}
