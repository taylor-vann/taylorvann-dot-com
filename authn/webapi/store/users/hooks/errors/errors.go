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
	FailedToCreateUser 	 = "failed to create User"
	FailedToDeleteUser	 = "failed to update User"
	FailedToDeleteUser 	 = "failed to delete User"
	FailedToReviveUser 	 = "failed to revive User"
	FailedToIndexUsers 	 = "failed to index Users"
	FailedToSearchUsers  = "failed to search Users"
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
