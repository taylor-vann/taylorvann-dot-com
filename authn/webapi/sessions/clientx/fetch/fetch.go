package fetch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	// "strconv"

	"log"

	"webapi/sessions/clientx/fetch/requests"
	"webapi/sessions/clientx/fetch/responses"
)

const (
	UsersStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	RolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	SessionsQueryAddress = "https://authn.briantaylorvann.com/q/sessions/"
	SessionsMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
)

var (
	Environemnt = os.Getenv("STAGE")

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

// create guest session

func CreateGuestSession(p *requests.GuestSession) (*string, error) {
	log.Println("FETCH CreateGuestSession")
	log.Println(p)
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
		return nil, errors.New("errors were returned in fetch")
	}

	return  &responseBody.Session.Token, nil
}


func ValidateGuestSession(p *requests.ValidateSession, sessionCookie *http.Cookie) (*string, error) {
	log.Println("FETCH ValidateGuestSession")
	log.Println(p)
	log.Println(sessionCookie)
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
		return nil, errors.New("errors were returned in fetch")
	}

	return  &responseBody.Session.Token, nil
}

func ValidateSession(p *requests.ValidateSession, sessionCookie *http.Cookie) (*string, error) {
	log.Println("FETCH ValidateSession")
	log.Println(p)
	log.Println(sessionCookie)
	
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
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.SessionBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return nil, errJson
	}
	if responseBody.Errors != nil {
		return nil, errors.New("errors were returned in fetch")
	}

	return  &responseBody.Session.Token, nil
}

func ValidateGuestUser(p *requests.ValidateGuestUser, sessionCookie *http.Cookie) (*responses.User, error) {
	log.Println("FETCH ValidateGuestUser")
	log.Println(p)
	log.Println(sessionCookie)
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
		return nil, errors.New("errors were returned in fetch")
	}

	users := *responseBody.Users
	if users != nil && len(users) > 0 {
		return &users[0], nil
	}

	return  nil, errors.New("nil session returned")
}

func ValidateInfraRole(p *requests.ValidateInfraRole, sessionCookie *http.Cookie) (*responses.Role, error) {
	log.Println("FETCH ValidateInfraRole called")
	log.Println(p)
	log.Println(sessionCookie)

	requestBodyBuffer, errRequestBodyBuffer := getRequestBodyBuffer(
		requests.Body{
			Action: "VALIDATE_INFRA_OVERLORD_ROLE",
			Params: p,
		},
	)
	if errRequestBodyBuffer != nil {
		log.Println(errRequestBodyBuffer)
		return nil, errRequestBodyBuffer
	}

	log.Println("about to create a request")

	req, errReq := http.NewRequest(
		"POST",
		RolesStoreQueryAddress,
		requestBodyBuffer,
	)
	if errReq != nil {
		log.Println(errReq)
		return nil, errReq
	}
	if sessionCookie == nil {
		log.Println("session cookie in validate infra role is nil")
		return nil, errors.New("session cookie is nil")
	}
	req.AddCookie(sessionCookie)

	resp, errResp := client.Do(req)
	if errResp != nil {
		log.Println(errResp)
		return nil, errResp
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("bad status code")
		return nil, errors.New(string(resp.StatusCode))
	}

	var responseBody responses.RolesBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		log.Println(errJson)
		return nil, errJson
	}
	if responseBody.Errors != nil {
		log.Println("nil respoonse, errors found")

		return nil, errors.New("errors were returned in fetch")
	}

	roles := *responseBody.Roles
	if roles != nil && len(roles) > 0 {
		return &roles[0], nil
	}
	
	return nil, errors.New("unable to validate infra role")
}

// create infram session

func CreateInfraSession(p *requests.InfraSession, guestSessionCookie *http.Cookie) (*string, error) {
	log.Println("FETCH CreateInfraSession:")
	log.Println(p)
	log.Println(guestSessionCookie)
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
		return nil, errors.New("errors were returned in fetch")
	}

	return  &responseBody.Session.Token, nil
}