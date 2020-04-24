// brian taylor vann
// taylorvann-dot-com

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
	CreateRole					       	= "CREATE_ROLE"
	ReadRole										= "READ_ROLE"
	SearchRoles									= "SEARCH_ROLES"
	IndexRoles									= "INDEX_ROLES"
	UpdateRole									= "UPDATE_ROLE"
	UpdateAccess								= "UPDATE_ACCESS"
	UpdateRole									= "UPDATE_ROLE"
	DeleteRole									= "DELETE_ROLE"
	UndeleteRole								= "UNDELETE_ROLE"
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
	case Read:
		queries.Read(w, &body)
	case Search:
		queries.Search(w, &body)
	case Index:
		queries.Index(w, &body)
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

	// switch body.Action {
	// case Create:
	// 	mutations.Read(w, &body)
	// case UpdateAccess:
	// 	mutations.UpdateAccess(w, &body)
	// case Update:
	// 	mutations.Update(w, &body)
	// case Delete:
	// 	mutations.Delete(w, &body)
	// case Undelete:
	// 	mutations.Undelete(w, &body)
	// default:
	// 	errors.CustomErrorResponse(w, errors.UnrecognizedMutation)
	}
}


