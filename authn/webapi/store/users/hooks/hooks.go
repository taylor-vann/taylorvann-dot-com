// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"encoding/json"
	"net/http"

	"log"
	
	"webapi/store/infrax/fetch"

	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	"webapi/store/users/hooks/mutations"
	"webapi/store/users/hooks/queries"

	"github.com/taylor-vann/tvgtb/jwtx"
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

func checkInfraSession(r *http.Request) (bool, error) {
	log.Println("validating infra session!")

	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		return false, errCookie
	}

	details, errDetails := jwtx.RetrieveTokenDetailsFromString(cookie.Value)
	if errDetails != nil {
		return false, errDetails
	}
	// log.Println(details.Payload)
	if details == nil {
		return false, nil
	}
	if details.Payload.Sub == "infra" && details.Payload.Aud == "public" {
		return true, nil
	}

	return false, nil
}

func validateInfraSession(r *http.Request) (bool, error) {
	sessionIsValid, errSessionIsValid := checkInfraSession(r)
	if !sessionIsValid || errSessionIsValid != nil {
		return sessionIsValid, errSessionIsValid
	}

	// remote check

}

func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
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
		queries.Read(w, &body)
	case ValidateGuest:
		queries.ValidateGuest(w, cookie, &body) // requires guest sessoion
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

	sessionIsValid, errSessionIsValid := checkInfraSession(r)
	if errSessionIsValid != nil {
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.UnrecognizedQuery,
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


