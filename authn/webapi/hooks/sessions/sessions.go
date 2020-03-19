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
	CreateGuestSession               = "CREATE_GUEST_SESSION"
	CreateGuestDocumentSession			 = "CREATE_GUEST_DOCUMENT_SESSION"
	CreatePublicSession              = "CREATE_PUBLIC_SESSION"
	CreatePublicDocumentSession			 = "CREATE_PUBLIC_DOCUMENT_SESSION"
	CreatePublicPasswordResetSession = "CREATE_PUBLIC_PASSWORD_RESET_SESSION"
	UpdateSession                    = "UPDATE_SESSION"
	ValidateSession                  = "VALIDATE_SESSION"
	RemoveSession                    = "REMOVE_SESSION"
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
	case CreateGuestSession:
		mutations.CreateGuestSession(w)
	case CreateGuestDocumentSession:
		mutations.CreateGuestDocumentSession(w)
	case CreatePublicSession:
		mutations.CreatePublicSession(w, &body)
	case CreatePublicDocumentSession:
		mutations.CreatePublicDocumentSession(w, &body)
	case CreatePublicPasswordResetSession:
		mutations.CreatePublicPasswordResetSession(w, &body)
	case UpdateSession:
		mutations.UpdateSession(w, &body)
	case RemoveSession:
		mutations.RemoveSession(w, &body)
	default:
		errors.CustomErrorResponse(w, errors.UnrecognizedMutation)
	}
}
