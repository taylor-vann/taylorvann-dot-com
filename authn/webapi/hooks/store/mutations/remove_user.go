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

// //	RemoveUsersRequestParams -
// type RemoveUserRequestParams struct {
// 	Email    string
// 	Password string
// }

// // RemoveUserRequestBody -
// type RemoveUserRequestBody struct {
// 	Action string                  `json:"action"`
// 	Params RemoveUserRequestParams `json:"params"`
// }

// // RemoveUser - must have guest credentials
// func RemoveUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("RemoveUser()")
// 	// session, errSession := validatePublicHeaders(r)
// 	// if errSession != nil {
// 	// 	defaultErrorResponse(w, errSession)
// 	// }
// 	// if session == nil {
// 	// 	errors.BadRequest(w, &errors.Payload{
// 	// 		Headers: &errors.InvalidHeadersProvided,
// 	// 	})
// 	// 	return
// 	// }

// 	var body RemoveUserRequestBody
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
// 	_, errUser := store.RemoveUser(&store.RemoveUserParams{
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
