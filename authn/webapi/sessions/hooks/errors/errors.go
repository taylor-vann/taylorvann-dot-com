package errors

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/responses"
)

const (
	ContentType = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	BadRequestFail 								 = "unable to decode request body"
	NilRequestBodyFail 						 = "request body is nil"
	UnrecognizedQuery 						 = "unrecognized query action requested"
	UnrecognizedMutation 					 = "unrecognized mutation action requested"
	UnrecognizedParams 						 = "unrecognized parameters"
	UnableToValidateSession 			 = "unable to create validate session"
	UnableToCreatePublicSession 	 = "unable to create public session"
	InvalidSessionCredentials 		 = "invalid session credentials provided"
	InvalidInfraCredentials 		 	 = "invalid infra session credentials provided"
	InvalidGuestCredentials 		 	 = "invalid guest session credentials provided"
	NilInfraCredentials 		 	 		 = "nil infra session credentials provided"
	InvalidSessionProvided    		 = "invalid session provided"
	CreateGuestSessionErrorMessage = "error creating guest session"
	InvalidRequestProvided         = "invalid request provided"
	InvalidDefaultUserProvided     = "invalid default user provided"
	UnableToMarshalSession         = "unable to marshal session"
	UnableToUpdateSession					 = "unable to update session"

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
