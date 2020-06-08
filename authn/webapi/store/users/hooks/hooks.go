// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"encoding/json"
	"net/http"

	// "log"

	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	"webapi/store/users/hooks/mutations"
	"webapi/store/users/hooks/queries"
)

const (
	Create				 = "CREATE_USER"
	Read					 = "READ_USER"
	ValidateGuest	 = "VALIDATE_GUEST_USER"
	Search				 = "SEARCH_USERS"
	Index					 = "INDEX_USERS"
	Update				 = "UPDATE_USER"
	UpdateEmail		 = "UPDATE_USER_EMAIL"
	UpdatePassword = "UPDATE_USER_PASSWORD"
	Delete				 = "DELETE_USER"
	Undelete			 = "UNDELETE_USER"

	SessionCookieHeader = "briantaylorvann.com_session"
)

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
}

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
	case ValidateGuest:
		queries.ValidateGuest(w, cookie, &body) // requires guest sessoion
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
	case UpdateEmail:
		mutations.UpdateEmail(w, cookie, &body)
	case UpdatePassword:
		mutations.UpdatePassword(w, cookie, &body)
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


