package fetches

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"webapi/sessionrequests/requests"
	"webapi/sessionrequests/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const (
	Post = "POST"

	UsersStoreQueryAddress  = "https://authn.briantaylorvann.com/q/users/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	ValidateUserAction        = "VALIDATE_USER"
	CreateClientSessionAction = "CREATE_CLIENT_SESSION"
	RemoveSessionAction       = "DELETE_SESSION"

	SessionCookieHeader = "briantaylorvann.com_session"
	cookieDomain        = ".briantaylorvann.com"
	ThreeDaysInSeconds  = 60 * 60 * 24 * 3
)

var (
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

func ValidateUser(params *requests.CreateSessionParams) (*responses.SafeRow, error) {
	if sessionx.InfraSession == nil {
		return nil, errNilInfraSession
	}

	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.ValidateUserBody{
			Action: ValidateUserAction,
			Params: params,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errBadParameters
	}

	req, errReq := http.NewRequest(
		Post,
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
		return nil, errUnsuccessfulSessionFetch
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsInFetchReturned
	}

	var user responses.SafeRow
	users := *responseBody.Users
	if len(users) > 0 {
		user = users[0]
	}

	return &user, nil
}

func GetClientSession(environment string, userID int64) (*http.Cookie, error) {
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.RequestSessionBody{
			Action: CreateClientSessionAction,
			Params: &requests.User{
				Environment: environment,
				UserID:      userID,
			},
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errBadParameters
	}

	req, errReq := http.NewRequest(
		Post,
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
		return nil, errUnsuccessfulSessionFetch
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsInFetchReturned
	}

	sessionCookie := createSessionCookie(responseBody.Session.Token)
	return sessionCookie, nil
}

func RemoveSession(
	environment string,
	sessionCookie *http.Cookie,
) error {
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		sessionCookie.Value,
	)
	if errTokenDetails != nil {
		return errTokenDetails
	}

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
		return errBadParameters
	}

	req, errReq := http.NewRequest(
		Post,
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return errReq
	}
	req.AddCookie(sessionx.InfraSession)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return errResp
	}
	if resp.StatusCode != http.StatusOK {
		return errUnsuccessfulSessionFetch
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return errJson
	}
	if responseBody.Errors != nil {
		return errErrorsInFetchReturned
	}

	return nil
}
