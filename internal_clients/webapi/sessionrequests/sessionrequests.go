package sessionrequests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"log"

	sessionErrors "webapi/sessionrequests/errors"
	"webapi/sessionrequests/requests"
	"webapi/sessionrequests/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/verifyx"
	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const (
	UsersStoreQueryAddress  = "https://authn.briantaylorvann.com/q/users/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	SessionCookieHeader = "briantaylorvann.com_session"
	sameSite            = "SameSite"
	cookieDomain        = "briantaylorvann.com"
	ThreeDaysInSeconds  = 60 * 60 * 24 * 3

	ValidateUser        = "VALIDATE_USER"
	CreateClientSession = "CREATE_CLIENT_SESSION"
	RemoveSessionAction = "DELETE_SESSION"
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

	NilSessionCookie = errors.New("nil session cookie provided")
)

func createSessionCookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:     SessionCookieHeader,
		Value:    session,
		MaxAge:   ThreeDaysInSeconds,
		Domain:   cookieDomain,
		Secure:   true,
		HttpOnly: true,
		SameSite: 3,
	}
}

func getRequestBodyBuffer(item interface{}) (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(item)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

// validate user, get safe user. return *user, error
func validateUser(params *requests.CreateSessionParams) (*responses.SafeRow, error) {
	if sessionx.InfraSession == nil {
		return nil, errors.New("infra session is nil")
	}

	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.ValidateUserBody{
			Action: ValidateUser,
			Params: params,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errors.New("bad parameters")
	}

	req, errReq := http.NewRequest(
		"POST",
		UsersStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionx.InfraSession)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful validation")
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errors.New("errors were returned in fetch")
	}

	var user responses.SafeRow
	users := *responseBody.Users
	if len(users) > 0 {
		user = users[0]
	}

	return &user, nil
}

// get client session, return *cookie, error
func fetchSession(environment string, userID int64) (*http.Cookie, error) {
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.RequestSessionBody{
			Action: CreateClientSession,
			Params: &requests.User{
				Environment: environment,
				UserID:      userID,
			},
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errors.New("bad parameters")
	}

	req, errReq := http.NewRequest(
		"POST",
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionx.InfraSession)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unsuccessful session fetch")
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errors.New("errors were returned in fetch")
	}

	sessionCookie := createSessionCookie(responseBody.Session.Token)
	return sessionCookie, nil
}

// get client session, return *cookie, error
func fetchRemoveSession(
	environment string,
	sessionCookie *http.Cookie,
) error {
	log.Println(sessionCookie)
	log.Println("fetch remove session")
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		sessionCookie.Value,
	)
	if errTokenDetails != nil {
		log.Println("error in token")
		log.Println(errTokenDetails)
		return errTokenDetails
	}

	log.Println("made it past sessions")
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.RemoveSessionBody{
			Action: RemoveSessionAction,
			Params: &requests.RemoveSessionParams{
				Environment: environment,
				Signature:   tokenDetails.Signature,
			},
		},
	)
	if errRequestBodyBuffer != nil {
		return errors.New("bad parameters")
	}

	req, errReq := http.NewRequest(
		"POST",
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return errReq
	}
	req.AddCookie(sessionx.InfraSession)

	log.Println("request created")

	resp, errResp := client.Do(req)
	if errResp != nil {
		return errResp
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("unsuccessful session fetch")
	}

	log.Println("finished request")

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		log.Println(errJson)
		return errJson
	}
	log.Println(responseBody)
	if responseBody.Errors != nil {
		log.Println("finished request with errors")
		log.Println(responseBody.Errors)
		return errors.New("errors were returned in fetch")
	}

	log.Println("finished request without errors")

	return nil
}

func addCookieAndReturnRequest(w http.ResponseWriter, sessionCookie *http.Cookie) {
	http.SetCookie(w, sessionCookie)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&responses.Body{})
}

// add client session to response
func RequestSession(w http.ResponseWriter, r *http.Request) {
	var params requests.CreateSessionParams
	errJsonDecode := json.NewDecoder(r.Body).Decode(&params)
	if errJsonDecode != nil {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	user, errUser := validateUser(&params)
	if errUser != nil {
		sessionErrors.DefaultResponse(w, errUser)
		return
	}

	sessionCookie, errSessionCookie := fetchSession(
		params.Environment,
		user.ID,
	)
	if errSessionCookie != nil {
		sessionErrors.DefaultResponse(w, errSessionCookie)
		return
	}

	addCookieAndReturnRequest(w, sessionCookie)
}

// remove session function
// write cookie as expired
func RemoveSession(w http.ResponseWriter, r *http.Request) {
	var params requests.RemoveSessionRequestParams
	errJsonDecode := json.NewDecoder(r.Body).Decode(&params)
	if errJsonDecode != nil {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	log.Println("made it to cookie")
	cookie, errCookie := r.Cookie(SessionCookieHeader)
	if errCookie != nil {
		sessionErrors.DefaultResponse(w, NilSessionCookie)
		return
	}

	if !verifyx.IsSessionValid(w, &verifyx.IsSessionValidParams{
		Environment:        params.Environment,
		InfraSessionCookie: sessionx.InfraSession,
		SessionCookie:      cookie,
	}) {
		sessionErrors.DefaultResponse(w, errJsonDecode)
		return
	}

	log.Println("verified session")

	errRemoveSession := fetchRemoveSession(params.Environment, cookie)
	if errRemoveSession != nil {
		log.Println("error verifying session")

		sessionErrors.DefaultResponse(w, errRemoveSession)
		return
	}

	addCookieAndReturnRequest(w, deletedCookie)
}
