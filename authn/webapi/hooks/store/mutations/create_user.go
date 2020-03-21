package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/store/errors"
	"webapi/store"
)

type CreateUserParams struct {
	User		store.CreateUserParams	`json:"user"`
	Session SessionParams					`json:"session"` 
}

// CreateUserRequestBody -
type CreateUserRequestBody struct {
	Action string           `json:"action"`
	Params CreateUserParams	`json:"params"`
}

// CreateUser - must have guest credentials
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// give it a session
	var body CreateUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		defaultErrorResponse(w, err)
		return
	}

	// create user
	user, errUser := store.CreateUser(&body.Params.User)
	if errUser != nil {
		errAsStr := errUser.Error()
		errors.BadRequest(w, &errors.Payload{
			Email:   &errors.UserAlreadyExists,
			Default: &errAsStr,
		})
		defaultErrorResponse(w, errUser)
		return
	}
	if user == nil {
		errAsStr := errUser.Error()
		errors.BadRequest(w, &errors.Payload{
			Email:   &errors.UserAlreadyExists,
			Default: &errAsStr,
		})
		defaultErrorResponse(w, errUser)
	}
}
