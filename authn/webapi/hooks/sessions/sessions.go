package sessions

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/mutations"
	"webapi/hooks/sessions/queries"
)

type ReadSessionAction struct {
	SessionSignature string `json:"session_signature"`
}
type RemoveSessionAction = ReadSessionAction
type MutationRequestPayload = mutations.RequestPayload
type MutationRequestBody = mutations.RequestBody
type MutationResponseBody = mutations.ResponseBody
type QueryRequestBody = queries.RequestBody
type QueryRequestPayload = queries.RequestPayload
type ErrorsPayload = mutations.ErrorsPayload

// Actions
const (
	CreateDocumentSession      = "CREATE_DOCUMENT_SESSION"
	CreateGuestSession         = "CREATE_GUEST_SESSION"
	CreatePublicSession        = "CREATE_PUBLIC_SESSION"
	CreateCreateAccountSession = "CREATE_CREATE_ACCOUNT_SESSION"
	CreateResetPasswordSession = "CREATE_RESET_PASSWORD_SESSION"
	CreateUpdateEmailSession   = "CREATE_UPDATE_EMAIL_SESSION"
	UpdateSession              = "UPDATE_SESSION"
	ValidateSession            = "VALIDATE_SESSION"
	RemoveSession              = "REMOVE_SESSION"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
		return
	}

	var body queries.RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errAsStr := err.Error()
		errors.BadRequest(w, &ErrorsPayload{
			Session: &errors.BadBodyFail,
			Default: &errAsStr,
		})
		return
	}

	switch body.Action {
	case ValidateSession:
		queries.ValidateSession(w, &body)
	default:
		errors.BadRequest(w, &ErrorsPayload{
			Session: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
		return
	}

	var body MutationRequestBody
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
	case CreateResetPasswordSession:
		mutations.CreateResetPasswordSession(w, &body)
	case CreateUpdateEmailSession:
		mutations.CreateUpdateEmailSession(w, &body)
	case UpdateSession:
		mutations.UpdateSession(w, &body)
	case RemoveSession:
		mutations.RemoveSession(w, &body)
	default:
		errors.CustomErrorResponse(w, errors.UnrecognizedMutation)
	}
}
