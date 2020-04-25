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
	Search				= "SEARCH_ROLES"
	Index					= "INDEX_ROLES"
	Update				= "UPDATE_ROLE"
	UpdateAccess	= "UPDATE_ACCESS"
	Delete				= "DELETE_ROLE"
	Undelete			= "UNDELETE_ROLE"
)

func Query(w http.ResponseWriter, r *http.Request) {
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
	case Read:
		queries.Read(w, &body)
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


