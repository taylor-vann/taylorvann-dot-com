// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	"webapi/store/users/hooks/mutations"
	"webapi/store/users/hooks/queries"
)

const (
	Create				 = "CREATE_USER"
	Read					 = "READ_USER"
	Validate			 = "VALIDATE_USER"
	Search				 = "SEARCH_USERS"
	Index					 = "INDEX_USERS"
	Update				 = "UPDATE_USER"
	UpdateEmail		 = "UPDATE_USER_EMAIL"
	UpdatePassword = "UPDATE_USER_PASSWORD"
	Delete				 = "DELETE_USER"
	Undelete			 = "UNDELETE_USER"
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
	case Validate:
		queries.Validate(w, &body)
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
	case UpdateEmail:
		mutations.UpdateEmail(w, &body)
	case UpdatePassword:
		mutations.UpdatePassword(w, &body)
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


