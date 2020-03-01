package store

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"webapi/hooks/store/errors"
	"webapi/hooks/store/mutations"
	"webapi/hooks/store/queries"
)

// RequestBodyParams -
type RequestBodyParams struct {
	Action string      `json:"action"`
	Params interface{} `json:"params`
}

// ResponseResult -
type ResponseResult struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

// ResponseErrors -
type ResponseErrors struct {
	User    *string `json:"user"`
	Email   *string `json:"email"`
	Default *string `json:"default"`
}

// ResponseBody -
type ResponseBody struct {
	Results *ResponseResult `json:"result"`
	Errors  *ResponseErrors `json:"errors"`
}

// Actions
const (
	// CreateUser -
	CreateUser = "CREATE_USER"
	// ReadUser -
	ReadUser = "READ_USER"
	// UpdateUserEmail -
	UpdateUserEmail = "UPDATE_USER_EMAIL"
	// UpdateUserPassword -
	UpdateUserPassword = "UPDATE_USER_PASSWORD"
	// RemoveUser -
	RemoveUser = "REMOVE_USER"
	// ReviveUser -
	ReviveUser = "REVIVE_USER"
)

func defaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	errors.BadRequest(w, &errors.Response{
		Default: &errAsStr,
	})
}

// Query - read information from session whitelist
func Query(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Response{
			Body: &errors.BadBodyFail,
		})
		return
	}

	var body RequestBodyParams
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		defaultErrorResponse(w, err)
		return
	}

	switch body.Action {
	case ReadUser:
	default:
		errors.BadRequest(w, &errors.Response{
			Store: &errors.UnrecognizedQuery,
		})
	}
}

// Mutation - mutate session whitelist
func Mutation(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &errors.Response{
			Body: &errors.BadBodyFail,
		})
		return
	}

	var body RequestBodyParams
	if r.Body != nil {
		b, _ := ioutil.ReadAll(r.Body)

		err := json.NewDecoder(
			ioutil.NopCloser(bytes.NewReader(b)),
		).Decode(&body)

		if err != nil {
			errors.BadRequest(w, &errors.Response{
				Default: &errors.BadBodyFail,
			})
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	switch body.Action {
	case CreateUser:
		mutations.CreateUser(w, r)
	case ReadUser:
		queries.ReadUser()
	case UpdateUserEmail:
		mutations.UpdateEmail(w, r)
	case UpdateUserPassword:
		mutations.UpdatePassword(w, r)
	case RemoveUser:
		mutations.RemoveUser(w, r)
	case ReviveUser:
		mutations.ReviveUser(w, r)
	default:
		errors.BadRequest(w, &errors.Response{
			Store: &errors.UnrecognizedMutation,
		})
	}
}
