// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"encoding/json"
	"net/http"

	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
	"webapi/store/roles/hooks/mutations"
	"webapi/store/roles/hooks/queries"
)

const (
	Create				= "CREATE_ROLE"
	Read					= "READ_ROLE"
	ValidateInfra = "VALIDATE_INFRA_OVERLORD_ROLE"
	Search				= "SEARCH_ROLES"
	Index					= "INDEX_ROLES"
	Update				= "UPDATE_ROLE"
	UpdateAccess	= "UPDATE_ROLE_ACCESS"
	Delete				= "DELETE_ROLE"
	Undelete			= "UNDELETE_ROLE"

	SessionCookieHeader = "briantaylorvann.com_session"
)

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.BadRequestFail,
		})
		return
	}

	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		errors.DefaultResponse(w, errCookie)
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.DefaultResponse(w, errJsonDecode)
		return
	}

	switch body.Action {
	case Read:
		queries.Read(w, cookie, &body)
	case ValidateInfra:
		queries.ValidateInfra(w, cookie, &body)
	case Search:
		queries.Search(w, cookie, &body)
	case Index:
		queries.Index(w, cookie, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.BadRequestFail,
		})
		return
	}

	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		errors.DefaultResponse(w, errCookie)
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.DefaultResponse(w, errJsonDecode)
		return
	}

	switch body.Action {
	case Create:
		mutations.Create(w, cookie, &body)
	case Update:
		mutations.Update(w, cookie, &body)
	case UpdateAccess:
		mutations.UpdateAccess(w, cookie, &body)
	case Delete:
		mutations.Delete(w, cookie, &body)
	case Undelete:
		mutations.Undelete(w, cookie, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.UnrecognizedMutation,
		})	
	}
}


