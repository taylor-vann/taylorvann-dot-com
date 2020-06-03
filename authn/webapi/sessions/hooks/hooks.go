package hooks

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/hooks/mutations"
	"webapi/sessions/hooks/queries"
)

const (
	CreateDocumentSession     	= "CREATE_DOCUMENT_SESSION"
	CreateGuestSession        	= "CREATE_GUEST_SESSION"
	CreateInfraOverlordSession  = "CREATE_INFRA_OVERLORD_SESSION"

	ValidateGuestSession        = "VALIDATE_GUEST_SESSION"
	CreatePublicSession       	= "CREATE_PUBLIC_SESSION"

	CreateCreateAccountSession	= "CREATE_CREATE_ACCOUNT_SESSION"
	CreateUpdatePasswordSession	= "CREATE_UPDATE_PASSWORD_SESSION"
	CreateUpdateEmailSession  	= "CREATE_UPDATE_EMAIL_SESSION"
	UpdateSession             	= "UPDATE_SESSION"
	ValidateSession           	= "VALIDATE_SESSION"
	DeleteSession             	= "DELETE_SESSION"

	SessionCookieHeader					= "briantaylorvann.com_session"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadRequestFail)
		return
	}
	
	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		errAsStr := errCookie.Error()
		errors.BadRequest(w, &responses.Errors{
			Default: &errAsStr,
		})
		return
	}

	var body requests.Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errAsStr := err.Error()
		errors.BadRequest(w, &responses.Errors{
			Default: &errAsStr,
		})
		return
	}
	
	switch body.Action {
	case ValidateGuestSession:	// the only public query
		queries.ValidateGuestSession(w, cookie, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomErrorResponse(w, errors.BadRequestFail)
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.CustomErrorResponse(w, errors.BadRequestFail)
		return
	}

	// only allow no session if "create guest session"
	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil && body.Action != CreateGuestSession {
		

		errAsStr := errCookie.Error()
		errors.BadRequest(w, &responses.Errors{
			Default: &errAsStr,
		})
		return
	}
	
	switch body.Action {
	// case CreateDocumentSession:
	// 	mutations.CreateDocumentSession(w, &body)
	case CreateGuestSession:
		mutations.CreateGuestSession(w, &body)	// the only public mutation
	case CreateInfraOverlordSession:
		mutations.CreateInfraSession(w, cookie, &body)	// the only guest mutation
	// case CreatePublicSession:
	// 	mutations.CreatePublicSession(w, &body)
	// case CreateCreateAccountSession:
	// 	mutations.CreateCreateAccountSession(w, &body)
	// case CreateUpdatePasswordSession:
	// 	mutations.CreateUpdatePasswordSession(w, &body)
	// case CreateUpdateEmailSession:
	// 	mutations.CreateUpdateEmailSession(w, &body)
	// case UpdateSession:
	// 	mutations.UpdateSession(w, &body)
	// case DeleteSession:
	// 	mutations.DeleteSession(w, &body)
	default:
		errors.CustomErrorResponse(w, errors.UnrecognizedMutation)
	}
}
