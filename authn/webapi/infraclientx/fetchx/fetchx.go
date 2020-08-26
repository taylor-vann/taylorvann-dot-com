package fetchx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"net/http"
	"os"

	"webapi/infraclientx/fetchx/requests"
	"webapi/infraclientx/fetchx/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const (
	UsersStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	RolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	SessionsQueryAddress = "https://authn.briantaylorvann.com/q/sessions/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"

	SessionCookieHeader	= "briantaylorvann.com_session"
)

var (
	Environment = os.Getenv("STAGE")

	ErrorsReturnedInFetch = errors.New("errors were returned in fetch")
	NilSessionReturned = errors.New("nil session returned")

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
		Name: SessionCookieHeader,
		Value: sessionToken,
	}
}

func CreateGuestSession(p *requests.GuestSession) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "CREATE_GUEST_SESSION",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		errMessage := fmt.Sprintf("%d status code returned", resp.StatusCode)
		return nil, errors.New(errMessage)
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, ErrorsReturnedInFetch
	}

	sessionCookie := createCookieFromString(responseBody.Session.Token)

	return  sessionCookie, nil
}


func ValidateGuestSession(
	p *requests.ValidateSession,
	sessionCookie *http.Cookie,
) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_GUEST_SESSION",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		return nil, ErrorsReturnedInFetch
	}

	returnedSessionCookie := createCookieFromString(responseBody.Session.Token)
	
	return  returnedSessionCookie, nil
}

func ValidateSession(
	p *requests.ValidateSession,
	sessionCookie *http.Cookie,
) (*http.Cookie, error) {	
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_SESSION",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		return nil, ErrorsReturnedInFetch
	}

	returnedSessionCookie := createCookieFromString(responseBody.Session.Token)
	
	return  returnedSessionCookie, nil
}

func ValidateGuestUser(
	p *requests.ValidateGuestUser,
	sessionCookie *http.Cookie,
) (*responses.User, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_GUEST_USER",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		return nil, ErrorsReturnedInFetch
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return  nil, errors.New("nil session returned")
}

func ValidateInfraRole(
	p *requests.ValidateInfraRole,
	sessionCookie *http.Cookie,
) (*responses.Role, error) {
	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_INFRA_OVERLORD_ROLE",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	if sessionCookie == nil {
		return nil, NilSessionReturned
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
		return nil, ErrorsReturnedInFetch
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}
	
	return nil, errors.New("unable to validate infra role")
}

func CreateInfraSession(
	p *requests.InfraSession,
	guestSessionCookie *http.Cookie,
) (*http.Cookie, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "CREATE_INFRA_OVERLORD_SESSION",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		errMessage := fmt.Sprintf("%d status code returned", resp.StatusCode)
		return nil, errors.New(errMessage)
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, ErrorsReturnedInFetch
	}

	sessionCookie := createCookieFromString(responseBody.Session.Token)

	return  sessionCookie, nil
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
		return nil, errors.New("could not parse a userID")
	}

	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: "READ_ROLE",
			Params: &requests.ValidateRole{
				Environment: p.Environment,
				UserID: int64(userID),
				Organization: "AUTHN_ADMIN",
			},
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		return nil, errReq
	}
	if infraSessionCookie == nil {
		return nil, NilSessionReturned
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
		return nil, ErrorsReturnedInFetch
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}
	
	return nil, errors.New("unable to validate role")
}

func ValidateUser(
	p *requests.ValidateUser,
	infraSessionCookie *http.Cookie,
) (*responses.User, error) {
	var requestBodyBuffer, errRequestBodyBuffer = getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_USER",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		return nil, errRequestBodyBuffer
	}

	req, errReq := http.NewRequest(
		"POST",
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
		return nil, ErrorsReturnedInFetch
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return  nil, errors.New("nil session returned")
}
