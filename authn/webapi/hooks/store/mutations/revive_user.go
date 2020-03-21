package mutations

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"webapi/hooks/constants"
// 	// "webapi/sessions"
// 	"webapi/hooks/store/errors"
// 	// "webapi/sessions"
// 	"webapi/store"
// )

// //	ReviveUsersRequestParams -
// type ReviveUserRequestParams struct {
// 	Email    string
// 	Password string
// }

// // ReviveUserRequestBody -
// type ReviveUserRequestBody struct {
// 	Action string                  `json:"action"`
// 	Params ReviveUserRequestParams `json:"params"`
// }

// // ReviveUser - must have guest credentials
// func ReviveUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("ReviveUser()")
// 	session, errSession := validatePublicHeaders(r)
// 	if errSession != nil {
// 		defaultErrorResponse(w, errSession)
// 	}
// 	if session == nil {
// 		errors.BadRequest(w, &errors.Payload{
// 			Headers: &errors.InvalidHeadersProvided,
// 		})
// 		return
// 	}

// 	var body ReviveUserRequestBody
// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		defaultErrorResponse(w, err)
// 		return
// 	}

// 	// verify password
// 	userRow, errValidate := store.ValidateUser(&store.ValidateUserParams{
// 		Email:    body.Params.Email,
// 		Password: body.Params.Password,
// 	})
// 	if errValidate != nil {
// 		errAsStr := errValidate.Error()
// 		errors.BadRequest(w, &errors.Payload{
// 			Password: &errors.UnableToValidateUser,
// 			Default:  &errAsStr,
// 		})
// 		return
// 	}
// 	if userRow == nil {
// 		errors.BadRequest(w, &errors.Payload{
// 			Email: &errors.UserDoesNotExist,
// 		})
// 		return
// 	}

// 	// create user
// 	_, errUser := store.ReviveUser(&store.ReviveUserParams{
// 		Email: body.Params.Email,
// 	})
// 	if errUser != nil {
// 		errAsStr := errUser.Error()
// 		errors.BadRequest(w, &errors.Payload{
// 			Email:   &errors.UserAlreadyExists,
// 			Default: &errAsStr,
// 		})
// 		return
// 	}

// 	csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set(constants.SessionTokenHeader, session.SessionToken)
// 	w.Header().Set(constants.CsrfTokenHeader, csrfAsBase64)
// 	w.WriteHeader(http.StatusOK)
// }
