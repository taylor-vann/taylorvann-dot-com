package store

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"webapi/hooks/store/errors"
	"webapi/hooks/store/mutations"
)

// We need to handle both users and roles and in a cache
// action 

// SessionParams -
type SessionParams = errors.SessionParams

// RequestPayload -
type RequestPayload struct {
	User    interface{}   `json:"user"`
}

// RequestBody -
type RequestBody struct {
	Action string         `json:"action"`
	Params RequestPayload `json:"params`
}

// ResponseErrors -
type ResponseErrors struct {
	User    *string `json:"user"`
	Email   *string `json:"email"`
	Default *string `json:"default"`
}

// ResponseBody -
type ResponseBody = errors.ResponseBody

// Actions
const (
	CreateUser = "CREATE_USER"
	ReadUser = "READ_USER"
	RetrieveUsers = "RETRIEVE_USERS"
	UpdateUserEmail = "UPDATE_USER_EMAIL"
	UpdateUserPassword = "UPDATE_USER_PASSWORD"
	RemoveUser = "REMOVE_USER"
	ReviveUser = "REVIVE_USER"
	CreateRole = "CREATE_ROLE"
	ReadRole 	 = "READ_ROLE"
	RetrieveRoles = "RETRIEVE_ROLES"
	RemoveRole = "REMOVE_ROLE"
)

func defaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	errors.BadRequest(w, &errors.Payload{
		Default: &errAsStr,
	})
}

// Query - read information from session whitelist
func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Payload{
			Body: &errors.BadBodyFail,
		})
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		defaultErrorResponse(w, err)
		return
	}

	switch body.Action {
	case ReadUser:
	default:
		errors.BadRequest(w, &errors.Payload{
			Store: &errors.UnrecognizedQuery,
		})
	}
}

// Mutation - mutate session whitelist
func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Payload{
			Body: &errors.BadBodyFail,
		})
		return
	}

	// get request body
	var body RequestBody
	bodyBytes, errBodyBytes := ioutil.ReadAll(r.Body)
	if errBodyBytes != nil {
		errors.BadRequest(w, &errors.Payload{
			Default: &errors.BadBodyFail,
		})
		return
	}
	errJsonDecode := json.NewDecoder(
		ioutil.NopCloser(bytes.NewReader(bodyBytes)),
	).Decode(&body)
	if errJsonDecode != nil {
		errors.BadRequest(w, &errors.Payload{
			Default: &errors.BadBodyFail,
		})
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	switch body.Action {
	case CreateUser:
		// mutations.CreateUser(w, r, validSession)
		mutations.CreateUser(w, r)
	case UpdateUserEmail:
		// mutations.UpdateEmail(w, r)
	case UpdateUserPassword:
		// mutations.UpdatePassword(w, r)
	case RemoveUser:
		// mutations.RemoveUser(w, r)
	case ReviveUser:
		// mutations.ReviveUser(w, r)
	case CreateRole:
		// mutations.CreateRole(w, r)
	case RemoveRole:
		// mutations.CreateRole(w, r)
	default:
		errors.BadRequest(w, &errors.Payload{
			Store: &errors.UnrecognizedMutation,
		})
	}
}
