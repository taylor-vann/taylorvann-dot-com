// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"encoding/json"
	"net/http"
	
	"log"

	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
	"webapi/store/roles/hooks/mutations"
	"webapi/store/roles/hooks/queries"
)

// we need to fetch validation from user
// then check if user has INTERNAL_INFRA role

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
			Body: &errors.BadRequestFail,
		})
		return
	}

	// only allow no session if "create guest session"
	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		log.Println("didn't find our session cookie!")

		errAsStr := errCookie.Error()
		errors.BadRequest(w, &responses.Errors{
			Default: &errAsStr,
		})
		return
	}

	var body requests.Body
	errJsonDecode := json.NewDecoder(r.Body).Decode(&body)
	if errJsonDecode != nil {
		errors.DefaultResponse(w, errJsonDecode)
		return
	}

	log.Println("query action: ", body.Action)
	switch body.Action {
	case Read:
		queries.Read(w, &body)
	case ValidateInfra:
		queries.ValidateInfra(w, cookie, &body)
	case Search:
		queries.Search(w, &body)
	case Index:
		queries.Index(w, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.UnrecognizedQuery,
		})
	}
}

func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
		})
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
		mutations.Create(w, &body)
	case Update:
		mutations.Update(w, &body)
	case UpdateAccess:
		mutations.UpdateAccess(w, &body)
	case Delete:
		mutations.Delete(w, &body)
	case Undelete:
		mutations.Undelete(w, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.UnrecognizedMutation,
		})	
	}
}


