package hooks

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/mutations"
	"webapi/sessions/hooks/queries"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
)

const (
	SessionCookieHeader					= "briantaylorvann.com_session"

	ValidateGuestSession        = "VALIDATE_GUEST_SESSION"
	ValidateSession        			= "VALIDATE_SESSION"
	
	CreateClientSession     		= "CREATE_CLIENT_SESSION"
	CreateCreateAccountSession	= "CREATE_CREATE_ACCOUNT_SESSION"
	CreateGuestSession        	= "CREATE_GUEST_SESSION"
	CreateInfraOverlordSession  = "CREATE_INFRA_OVERLORD_SESSION"
	CreateUpdateEmailSession  	= "CREATE_UPDATE_EMAIL_SESSION"
	CreateUpdatePasswordSession	= "CREATE_UPDATE_PASSWORD_SESSION"
	DeleteSession             	= "DELETE_SESSION"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.CustomResponse(w, errors.NilRequestBodyFail)
		return
	}

	var body requests.Body
	errDecode := json.NewDecoder(r.Body).Decode(&body)
	if errDecode != nil {
		errors.DefaultResponse(w, errDecode)
		return
	}

	cookie, _ := r.Cookie(SessionCookieHeader)

	switch body.Action {
	case ValidateGuestSession:	// the only public guest query
		queries.ValidateGuestSession(w, &body)
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
		errors.CustomResponse(w, errors.NilRequestBodyFail)
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		
		errors.CustomResponse(w, errors.BadRequestFail)
		return
	}
	
	cookie, _ := r.Cookie(SessionCookieHeader)
	
	switch body.Action {
	case CreateGuestSession:  // the only public mutation
		mutations.CreateGuestSession(w, &body)
	case CreateInfraOverlordSession:  // the only guest mutation
		mutations.CreateInfraSession(w, cookie, &body)
	case CreateClientSession:
		mutations.CreateClientSession(w, cookie, &body)
	case CreateCreateAccountSession:
		mutations.CreateCreateAccountSession(w, cookie, &body)
	case CreateUpdatePasswordSession:
		mutations.CreateUpdatePasswordSession(w, cookie, &body)
	case CreateUpdateEmailSession:
		mutations.CreateUpdateEmailSession(w, cookie, &body)
	case DeleteSession:
		mutations.DeleteSession(w, cookie, &body)
	default:
		errors.CustomResponse(w, errors.UnrecognizedMutation)
	}
}
