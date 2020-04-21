package sessions

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/requests"
	"webapi/hooks/sessions/responses"
	"webapi/hooks/sessions/mutations"
	"webapi/hooks/sessions/queries"
)

const (
	CreateDocumentSession     	= "CREATE_DOCUMENT_SESSION"
	CreateGuestSession        	= "CREATE_GUEST_SESSION"
	CreatePublicSession       	= "CREATE_PUBLIC_SESSION"
	CreateCreateAccountSession	= "CREATE_CREATE_ACCOUNT_SESSION"
	CreateUpdatePasswordSession	= "CREATE_UPDATE_PASSWORD_SESSION"
	CreateUpdateEmailSession  	= "CREATE_UPDATE_EMAIL_SESSION"
	UpdateSession             	= "UPDATE_SESSION"
	ValidateSession           	= "VALIDATE_SESSION"
	DeleteSession             	= "DELETE_SESSION"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
		return
	}

	var body requests.Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errAsStr := err.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &errors.BadBodyFail,
			Default: &errAsStr,
		})
		return
	}

	switch body.Action {
	case ValidateSession:
		queries.ValidateSession(w, &body)
	default:
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
		return
	}

	switch body.Action {
	case CreateDocumentSession:
		mutations.CreateDocumentSession(w)
	case CreateGuestSession:
		mutations.CreateGuestSession(w)
	case CreatePublicSession:
		mutations.CreatePublicSession(w, &body)
	case CreateCreateAccountSession:
		mutations.CreateCreateAccountSession(w, &body)
	case CreateUpdatePasswordSession:
		mutations.CreateUpdatePasswordSession(w, &body)
	case CreateUpdateEmailSession:
		mutations.CreateUpdateEmailSession(w, &body)
	case UpdateSession:
		mutations.UpdateSession(w, &body)
	case DeleteSession:
		mutations.RemoveSession(w, &body)
	default:
		errors.CustomErrorResponse(w, errors.UnrecognizedMutation)
	}
}
