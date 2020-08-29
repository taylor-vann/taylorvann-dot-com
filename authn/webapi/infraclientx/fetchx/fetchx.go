package fetchx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
	"webapi/infraclientx/fetchx/requests"
	"webapi/infraclientx/fetchx/responses"
)

const (
	UsersStoreQueryAddress  = "https://authn.briantaylorvann.com/q/users/"
	RolesStoreQueryAddress  = "https://authn.briantaylorvann.com/q/roles/"
	SessionsQueryAddress    = "https://authn.briantaylorvann.com/q/sessions/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	SessionCookieHeader = "briantaylorvann.com_session"

	Post = "POST"

	CreateGuestSessionAction   = "CREATE_GUEST_SESSION"
	ValidateGuestSessionAction = "VALIDATE_GUEST_SESSION"
	ValidateSessionAction      = "VALIDATE_SESSION"
	ValidateGuestUserAction    = "VALIDATE_GUEST_USER"
	ValidateUserAction         = "VALIDATE_USER"
	ValidateInfraRoleAction    = "VALIDATE_INFRA_OVERLORD_ROLE"
	CreateInfraSessionAction   = "CREATE_INFRA_OVERLORD_SESSION"
	ReadRoleAction             = "READ_ROLE"

	AuthnAdmin = "AUTHN_ADMIN"
)

var (
	Environment = os.Getenv("STAGE")

	errErrorsReturnedInFetch = errors.New("errors were returned in fetch")
	errNilSessionReturned    = errors.New("nil session returned")
	errUnableToValidateUser  = errors.New("unable to validate user")
	errUnableToValidateRole  = errors.New("unable to validate role")
	errUnableToCreateSession = errors.New("unable to create session")
	errUnableToParseID       = errors.New("unable to parse a userID")

	client = http.Client{}
)

func getRequestBodyBuffer(item interface{}) (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(item)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func createCookieFromString(sessionToken string) *http.Cookie {
	return &http.Cookie{
		Name:  SessionCookieHeader,
		Value: sessionToken,
	}
}

func CreateGuestSession(p *requests.GuestSession) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: CreateGuestSessionAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errUnableToCreateSession
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	sessionCookie := createCookieFromString(responseBody.Session.Token)

	return sessionCookie, nil
}

func ValidateGuestSession(
	p *requests.ValidateSession,
	sessionCookie *http.Cookie,
) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: ValidateGuestSessionAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		SessionsQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	returnedSessionCookie := createCookieFromString(responseBody.Session.Token)

	return returnedSessionCookie, nil
}

func ValidateSession(
	p *requests.ValidateSession,
	sessionCookie *http.Cookie,
) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: ValidateSessionAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		SessionsQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode))
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	returnedSessionCookie := createCookieFromString(responseBody.Session.Token)

	return returnedSessionCookie, nil
}

func ValidateGuestUser(
	p *requests.ValidateGuestUser,
	sessionCookie *http.Cookie,
) (*responses.User, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: ValidateGuestUserAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		UsersStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return nil, errUnableToValidateUser
}

func ValidateInfraRole(
	p *requests.ValidateInfraRole,
	sessionCookie *http.Cookie,
) (*responses.Role, error) {
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: ValidateInfraRoleAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	if sessionCookie == nil {
		return nil, errNilSessionReturned
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.RolesBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}

	return nil, errUnableToValidateRole
}

func CreateInfraSession(
	p *requests.InfraSession,
	guestSessionCookie *http.Cookie,
) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: CreateInfraSessionAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		SessionsMutationAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(guestSessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errUnableToCreateSession
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	sessionCookie := createCookieFromString(responseBody.Session.Token)

	return sessionCookie, nil
}

func ValidateRoleFromSession(
	p *requests.ValidateRoleFromSession,
	infraSessionCookie *http.Cookie,
) (*responses.Role, error) {
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(p.Token)
	if errTokenDetails != nil {
		return nil, errTokenDetails
	}

	userID, errUserID := strconv.Atoi(tokenDetails.Payload.Aud)
	if errUserID != nil {
		return nil, errUnableToParseID
	}

	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: ReadRoleAction,
			Params: &requests.ValidateRole{
				Environment:  p.Environment,
				UserID:       int64(userID),
				Organization: AuthnAdmin,
			},
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	if infraSessionCookie == nil {
		return nil, errNilSessionReturned
	}
	req.AddCookie(infraSessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.RolesBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}

	return nil, errUnableToValidateRole
}

func ValidateUser(
	p *requests.ValidateUser,
	infraSessionCookie *http.Cookie,
) (*responses.User, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: ValidateUserAction,
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		Post,
		UsersStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	req.AddCookie(infraSessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errErrorsReturnedInFetch
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return nil, errUnableToValidateUser
}
