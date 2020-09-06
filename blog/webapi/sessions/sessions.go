package sessions

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"webapi/sessions/fetches"

	sessionErrors "webapi/sessions/errors"
	"webapi/sessions/requests"
	"webapi/sessions/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/verifyx"
)

const (
	UsersStoreQueryAddress  = "https://authn.briantaylorvann.com/q/users/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	SessionCookieHeader = "briantaylorvann.com_session"
	cookieDomain        = ".briantaylorvann.com"
	ThreeDaysInSeconds  = 60 * 60 * 24 * 3

	ValidateUser        = "VALIDATE_USER"
	CreateClientSession = "CREATE_CLIENT_SESSION"
	RemoveSessionAction = "DELETE_SESSION"

	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

var (
	Environment = os.Getenv("STAGE")

	client = http.Client{}

	deletedCookie = &http.Cookie{
		Name:     SessionCookieHeader,
		MaxAge:   -1,
		Value:    "",
		Domain:   cookieDomain,
		Secure:   true,
		HttpOnly: true,
		SameSite: 3,
	}

	errNilSessionCookie         = errors.New("nil session cookie provided")
	errNilInfraSession          = errors.New("infra session is nil")
	errBadParameters            = errors.New("bad parameters")
	errUnsuccessfulSessionFetch = errors.New("unsuccessful session fetch")
	errErrorsInFetchReturned    = errors.New("errors were returned in fetch")
)

func addJsonHeader(w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJson)
}

func addCookieAndReturnRequest(w http.ResponseWriter, sessionCookie *http.Cookie) {
	w.WriteHeader(http.StatusOK)
	http.SetCookie(w, sessionCookie)
	json.NewEncoder(w).Encode(&responses.Body{})
}

func RequestGuestSession(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)
	if sessionx.InfraSession == nil {
		sessionErrors.DefaultResponse(w, errNilInfraSession)
		return
	}

	var params requests.CreateGuestSessionParams
	errJsonDecode := json.NewDecoder(r.Body).Decode(&params)
	if errJsonDecode != nil {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	sessionCookie, errSessionCookie := fetches.GetGuestSession(&params)
	if errSessionCookie != nil {
		sessionErrors.DefaultResponse(w, errSessionCookie)
		return
	}

	addCookieAndReturnRequest(w, sessionCookie)
}

func RequestClientSession(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)
	if sessionx.InfraSession == nil {
		sessionErrors.DefaultResponse(w, errNilInfraSession)
		return
	}

	var params requests.CreateClientSessionParams
	errJsonDecode := json.NewDecoder(r.Body).Decode(&params)
	if errJsonDecode != nil {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	user, errUser := fetches.ValidateUser(&params)
	if errUser != nil {
		sessionErrors.DefaultResponse(w, errUser)
		return
	}

	sessionCookie, errSessionCookie := fetches.GetClientSession(
		params.Environment,
		user.ID,
	)
	if errSessionCookie != nil {
		sessionErrors.DefaultResponse(w, errSessionCookie)
		return
	}

	addCookieAndReturnRequest(w, sessionCookie)
}

func RemoveSession(w http.ResponseWriter, r *http.Request) {
	addJsonHeader(w)
	if sessionx.InfraSession == nil {
		sessionErrors.DefaultResponse(w, errNilInfraSession)
		return
	}

	var params requests.RemoveSessionParams
	errJsonDecode := json.NewDecoder(r.Body).Decode(&params)
	if errJsonDecode != nil {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		sessionErrors.DefaultResponse(w, errNilSessionCookie)
		return
	}

	if !verifyx.IsSessionValid(&verifyx.IsSessionValidParams{
		Environment:        params.Environment,
		InfraSessionCookie: sessionx.InfraSession,
		SessionCookie:      cookie,
	}) {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	errRemoveSession := fetches.RemoveSession(params.Environment, cookie)
	if errRemoveSession != nil {
		sessionErrors.DefaultResponse(w, errRemoveSession)
		return
	}

	addCookieAndReturnRequest(w, deletedCookie)
}