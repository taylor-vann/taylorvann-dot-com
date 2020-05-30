package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	// "net/url"
	"os"
	// "time"

	"log"

	"webapi/store/infrax/requests"
	"webapi/store/infrax/responses"

	"golang.org/x/net/publicsuffix"
)

const (
	ApplicationJson = "application/json"
	AuthnUserStoreQueryAddress = "https://authn.briantaylorvann.com/q/users/"
	AuthnRolesStoreQueryAddress = "https://authn.briantaylorvann.com/q/roles/"
	AuthnSessionMutationAddress = "https://authn.briantaylorvann.com/m/sessions/"
	CookieDomain = "www.briantaylorvann.com"
	GuestSessionExpirationInSeconds = 60 * 60 * 24 * 3
	SessionCookieHeader = "briantaylorvann.com_session"
)

var (
	Environemnt = os.Getenv("STAGE")
	InfraOverlordEmail = os.Getenv("INFRA_OVERLORD_EMAIL")
	InfraOverlordPassword = os.Getenv("INFRA_OVERLORD_PASSWORD")

	RequestGuestSessionBody = requests.Body{
		Action: "CREATE_GUEST_SESSION",
		Params: requests.GuestSessionParams {
			Environment: Environemnt,
		},
	}

	RequestValidateUserBody = requests.Body{
		Action: "VALIDATE_GUEST_USER",
		Params: requests.ValidateUserParams {
			Environment: Environemnt,
			Email: InfraOverlordEmail,
			Password: InfraOverlordPassword,
		},
	}
)

var CookieJar *cookiejar.Jar
var Client *http.Client

func getGuestSessionRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestGuestSessionBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

func getValidateUserRequestBodyBuffer() (*bytes.Buffer, error) {
	sessionBuffer := new(bytes.Buffer)
	errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestValidateUserBody)
	if errJsonBuffer != nil {
		return nil, errJsonBuffer
	}

	return sessionBuffer, nil
}

// func CreateSessionCookie(sessionToken string) *http.Cookie {
// 	return &http.Cookie{
// 		Name:			SessionCookieHeader,
// 		Value:		sessionToken,
// 		MaxAge:		GuestSessionExpirationInSeconds,
// 		Domain:   CookieDomain,
// 		Path:     "/",
// 		Secure:		true,
// 		HttpOnly:	true,
// 		SameSite:	3,
// 	}
// }

func FetchGuestSession() (string, error) {
	var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getGuestSessionRequestBodyBuffer()
	if errGuestSessionRequestBodyBuffer != nil {
		return "", errGuestSessionRequestBodyBuffer
	}


	resp, errResp := http.Post(
		AuthnSessionMutationAddress,
		ApplicationJson,
		guestSessionRequestBodyBuffer,
	)
	if errResp != nil {
		return "", errResp
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.Body
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		return "", errJson
	}
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}
	if responseBody.Session != nil {
		return responseBody.Session.Token, nil
	}

	return  "", errors.New("nil session returned")
}

// go for guest session to prove it first
func FetchValidateUser() (string, error) {
	var guestSessionRequestBodyBuffer, errGuestSessionRequestBodyBuffer = getValidateUserRequestBodyBuffer()
	if errGuestSessionRequestBodyBuffer != nil {
		return "", errGuestSessionRequestBodyBuffer
	}

	resp, errResp := Post(
		AuthnUserStoreQueryAddress,
		guestSessionRequestBodyBuffer,
	)
	if errResp != nil {
		log.Println(errResp.Error())
		return "", errResp
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(resp.StatusCode))
	}

	var responseBody responses.UsersBody
	errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJson != nil {
		log.Println(string(resp.StatusCode))
		log.Println(errJson)
		return "", errJson
	}
	if responseBody.Errors != nil {
			return "", errors.New("errors were returned in fetch")
	}
	if responseBody.Users != nil {
		return "found a user", nil
	}

	return  "", errors.New("nil session returned")
	// sessionBuffer := new(bytes.Buffer)
	// errJsonBuffer := json.NewEncoder(sessionBuffer).Encode(RequestValidateUserBody)
	// if errJsonBuffer != nil {
	// 	return nil, errJsonBuffer
	// }

	// // var validateUserRequestBodyBuffer, errValidateUserRequestBodyBuffer = getValidateUserRequestBodyBuffer()
	// // if errValidateUserRequestBodyBuffer != nil {
	// // 	return nil, errValidateUserRequestBodyBuffer
	// // }

	// resp, errResp := Post(
	// 	AuthnUserStoreQueryAddress,
	// 	sessionBuffer,
	// )
	// log.Println("posted request")

	// if errResp != nil {
	// 	log.Println("error in response")

	// 	log.Println(errResp.Error())
	// 	return nil, errResp
	// }
	
	// if resp.StatusCode != http.StatusOK {
	// 	log.Println("bad status code")

	// 	log.Println(string(resp.StatusCode))
	// 	log.Println(resp)
	// 	log.Println(resp.Body)
		
	// 	// return nil, errors.New(string(resp.StatusCode))
	// }
	// log.Println("finished validate user request")

	// var responseBody responses.UsersBody
	// errJson := json.NewDecoder(resp.Body).Decode(&responseBody)
	// if errJson != nil {
	// 	log.Println(string(resp.StatusCode))
	// 	log.Println(errJson)
	// 	return nil, errJson
	// }
	// if responseBody.Errors != nil {
	// 		log.Println(responseBody.Errors)
	// 		log.Println(responseBody.Users)


	// 		return nil, errors.New("errors were returned in fetch")
	// }
	// if responseBody.Users != nil  {
	// 	log.Println("made it to response")
	// 	users := *responseBody.Users
	// 	if len(users) > 0 {
	// 		// validatedUser := responseBody.Users[0]
	// 		// return &validatedUser, nil
	// 		validatedUser := users[0]
	// 		return &validatedUser, nil
	// 	}
	// }
	// log.Println("yoo about to return nil user")

	// return nil, errors.New("nil user returned")
}

func Init() (*http.Client, error) {
	cookiejar, errCookiejar := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	// cookiejar, errCookiejar := cookiejar.New(nil)
	if errCookiejar != nil {
		return nil, errCookiejar
	}
	CookieJar = cookiejar
	if cookiejar == nil {
		return nil, errors.New("nil jar provided")
	}
	
	
	Client = &http.Client{
		Jar: cookiejar,
		// Transport: &http.Transport{
		// 	MaxIdleConns:       10,
		// 	IdleConnTimeout:    30 * time.Second,
		// 	DisableCompression: true,
		// 	ForceAttemptHTTP2: 	true,
		// },
	}

	_, errSession := FetchGuestSession()
	if errSession != nil {
		return nil, errSession
	}

	// var cookies []*http.Cookie
	// cookies = append(cookies, CreateSessionCookie(session))
	// urlStr, errUrlStr := url.Parse(CookieDomain)
	// if errUrlStr != nil {
	// 	return nil, errUrlStr
	// }
	// cookiejar.SetCookies(urlStr, cookies)

	// get internal infra session

	// validate user

	return Client, nil
}

// validate user with session
// then change to validate internal role

// GetInternalSession()
// -> create internal session (1 year, something ridiculous)
// -> validate internal user endpoint // needs password & guest session

func Post(urlStr string, payload *bytes.Buffer) (*http.Response, error) {
	request, errRequest := http.NewRequest("POST", urlStr, payload)
	if errRequest != nil {
		log.Println("error making request")
		log.Println(errRequest)

		return nil, errRequest
	}

	request.Header.Add("Content-Type", "application/json")

	// try it with a post
	return Client.Do(request)
	// return Client.Post(urlStr, "application/json", payload)
	// return http.Post(urlStr, "application/json", payload)

	// return http.Post()
}