package errors

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/hooks/responses"
)

var (
	BadRequestFail  		 		 	 = "unable to decode request body"
	UnrecognizedQuery 	 		 	 = "unrecognized query action requested"
	UnrecognizedMutation 		 	 = "unrecognized mutation action requested"
	FailedToCreateUser 	 		 	 = "failed to create User"
	FailedToReadUser	 	 		 	 = "failed to read User"
	FailedToIndexUsers 	 		 	 = "failed to index Users"
	FailedToSearchUsers  		 	 = "failed to search Users"
	FailedToUpdateUser	 	 	 	 = "failed to Update User"
	FailedToUpdateEmailUser 	 = "failed to Update Email User"
	FailedToUpdatePasswordUser = "failed to Update Password User"
	FailedToDeleteUser 	 		 	 = "failed to delete User"
	FailedToUndeleteUser 		 	 = "failed to undelete User"
)

var defaultFail = "unable to return Roles"

func DefaultResponse(w http.ResponseWriter, err error) {
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
