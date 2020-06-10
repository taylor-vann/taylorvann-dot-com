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
	SessionCookieHeader					= "briantaylorvann.com_session"

	ValidateGuestSession        = "VALIDATE_GUEST_SESSION"
	ValidateSession        			= "VALIDATE_SESSION"

	CreateClientSession     		= "CREATE_DOCUMENT_SESSION"
	CreateGuestSession        	= "CREATE_GUEST_SESSION"
	CreateInfraOverlordSession  = "CREATE_INFRA_OVERLORD_SESSION"
	// CreateCreateAccountSession		= "CREATE_CREATE_ACCOUNT_SESSION"
	// CreateUpdatePasswordSession	= "CREATE_UPDATE_PASSWORD_SESSION"
	// CreateUpdateEmailSession  		= "CREATE_UPDATE_EMAIL_SESSION"
	DeleteSession             	= "DELETE_SESSION"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomResponse(w, errors.BadRequestFail)
		return
	}
	
	cookie, _ := r.Cookie(SessionCookieHeader)

	var body requests.Body
	errDecode := json.NewDecoder(r.Body).Decode(&body)
	if errDecode != nil {
		errors.DefaultResponse(w, errDecode)
		return
	}

	switch body.Action {
	case ValidateGuestSession:	// the only public guest query
		queries.ValidateGuestSession(w, cookie, &body)
	case ValidateSession:
		queries.ValidateSession(w, cookie, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomResponse(w, errors.BadRequestFail)
		return
	}

	cookie, _ := r.Cookie(SessionCookieHeader)

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.CustomResponse(w, errors.BadRequestFail)
		return
	}
	
	switch body.Action {
	case CreateGuestSession:
		mutations.CreateGuestSession(w, &body)	// the only public mutation
	case CreateInfraOverlordSession:
		mutations.CreateInfraSession(w, cookie, &body)	// the only guest mutation
	case CreateClientSession:
		mutations.CreateClientSession(w, cookie, &body)
	// case CreatePublicSession:
	// 	mutations.CreatePublicSession(w, &body)
	// case CreateCreateAccountSession:
	// 	mutations.CreateCreateAccountSession(w, &body)
	// case CreateUpdatePasswordSession:
	// 	mutations.CreateUpdatePasswordSession(w, &body)
	// case CreateUpdateEmailSession:
	// 	mutations.CreateUpdateEmailSession(w, &body)
	case DeleteSession:
		mutations.DeleteSession(w, cookie, &body)
	default:
		errors.CustomResponse(w, errors.UnrecognizedMutation)
	}
}
