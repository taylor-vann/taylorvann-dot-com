package mutations

import (
	"encoding/base64"
	"encoding/json"
	// "fmt"
	"net/http"

	"webapi/hooks/constants"
	// "webapi/sessions"
	"webapi/hooks/store/errors"
	// "webapi/sessions"
	"webapi/store"
)

//	UpdatePasswordsRequestParams -
type UpdatePasswordRequestParams struct {
	Email           string
	Password        string
	UpdatedPassword string
}

// UpdateUserPasswordRequestBody -
type UpdateUserPasswordRequestBody struct {
	Action string                      `json:"action"`
	Params UpdatePasswordRequestParams `json:"params"`
}

// UpdatePassword - must have guest credentials
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	session, errSession := validatePublicHeaders(r)
	if errSession != nil {
		defaultErrorResponse(w, errSession)
	}
	if session == nil {
		errors.BadRequest(w, &errors.Response{
			Headers: &errors.InvalidHeadersProvided,
		})
		return
	}

	var body UpdateUserPasswordRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		defaultErrorResponse(w, err)
		return
	}

	// verify password
	userRow, errValidate := store.ValidateUser(&store.ValidateUserParams{
		Email:    body.Params.Email,
		Password: body.Params.Password,
	})
	if errValidate != nil {
		errAsStr := errValidate.Error()
		errors.BadRequest(w, &errors.Response{
			Password: &errors.UnableToValidateUser,
			Default:  &errAsStr,
		})
		return
	}
	if userRow == nil {
		errors.BadRequest(w, &errors.Response{
			Email: &errors.UserDoesNotExist,
		})
		return
	}

	// create user
	_, errUser := store.UpdatePassword(&store.UpdatePasswordParams{
		Email:           body.Params.Email,
		UpdatedPassword: body.Params.UpdatedPassword,
	})
	if errUser != nil {
		errAsStr := errUser.Error()
		errors.BadRequest(w, &errors.Response{
			Email:   &errors.UserAlreadyExists,
			Default: &errAsStr,
		})
		return
	}

	csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(constants.SessionTokenHeader, session.SessionToken)
	w.Header().Set(constants.CsrfTokenHeader, csrfAsBase64)
	w.WriteHeader(http.StatusOK)
}
