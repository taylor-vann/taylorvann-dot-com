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

// // CreateUpdateUserEmailRequestBodyParams -
// type CreateUpdateUserEmailRequestBodyParams struct {
// 	Action string                  `json:"action"`
// 	Params store.UpdateEmailParams `json:"params"`
// }

// // UpdateEmail - must have guest credentials
// func UpdateEmail(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("made it here")
// 	fmt.Println(r.Header)
// 	session, errSession := validatePublicHeaders(r)
// 	if errSession != nil {
// 		defaultErrorResponse(w, errSession)
// 	}
// 	if session == nil {
// 		errors.BadRequest(w, &errors.Response{
// 			Headers: &errors.InvalidHeadersProvided,
// 		})
// 		return
// 	}

// 	var body CreateUpdateUserEmailRequestBodyParams
// 	err := json.NewDecoder(r.Body).Decode(&body)
// 	if err != nil {
// 		defaultErrorResponse(w, err)
// 		return
// 	}

// 	// create user
// 	_, errUser := store.UpdateEmail(&body.Params)
// 	if errUser != nil {
// 		errAsStr := errUser.Error()
// 		errors.BadRequest(w, &errors.Response{
// 			Email:   &errors.UserAlreadyExists,
// 			Default: &errAsStr,
// 		})
// 		defaultErrorResponse(w, errUser)
// 		return
// 	}

// 	csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set(constants.SessionTokenHeader, session.SessionToken)
// 	w.Header().Set(constants.CsrfTokenHeader, csrfAsBase64)
// 	w.WriteHeader(http.StatusOK)
// }
