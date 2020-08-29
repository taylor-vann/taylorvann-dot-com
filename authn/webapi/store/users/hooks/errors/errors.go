package errors

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/hooks/responses"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	BadRequestFail               = "unable to decode request body"
	NilRequestBodyFail           = "nil body provided in request"
	UnrecognizedQuery            = "unrecognized query action requested"
	UnrecognizedMutation         = "unrecognized mutation action requested"
	UnrecognizedParams           = "unrecognized params in request"
	InvalidSession               = "invalid session"
	InvalidGuestSession          = "invalid guest session"
	InvalidInfraSession          = "invalid infra session"
	NilInfraCredentials          = "nil infra session credentials provided"
	FailedToValidateGuestSession = "failed to validate guest session"
	FailedToCreateUser           = "failed to create User"
	FailedToReadUser             = "failed to read User"
	FailedToValidateUser         = "failed to validate user"
	FailedToIndexUsers           = "failed to index Users"
	FailedToSearchUsers          = "failed to search Users"
	FailedToUpdateUser           = "failed to Update User"
	FailedToUpdateEmailUser      = "failed to Update Email User"
	FailedToUpdatePasswordUser   = "failed to Update Password User"
	FailedToDeleteUser           = "failed to delete User"
	FailedToUndeleteUser         = "failed to undelete User"

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
			Default: &defaultFail,
		},
	})
}
