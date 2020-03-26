package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/sessions"
	"webapi/sessions/constants"
)

// 	trade a guest document session for a password reset session
func CreateResetPasswordSession(w http.ResponseWriter, requestBody *RequestBody) {
	if requestBody.Params == nil || requestBody.Params.Credentials == nil {
		errors.CustomErrorResponse(w, InvalidRequestProvided)
		return
	}

	validRequest, errValidRequest := validateAndRemoveSession(
		requestBody,
		constants.Document,
		constants.Guest,
	)
	if errValidRequest != nil {
		errAsStr := errValidRequest.Error()
		errors.BadRequest(w, &errors.ResponsePayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}
	if !validRequest {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
		return
	}

	session, errSession := sessions.Create(
		sessions.ComposeResetPasswordSessionParams(&sessions.CreateAccountParams{
			Email: *requestBody.Params.Credentials.Email,
		}),
	)

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(&ResponsePayload{
			SessionToken: &session.SessionToken,
		})
		if errMarshal == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&marshalledJSON)
			return
		}

		errors.CustomErrorResponse(w, UnableToMarshalSession)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &errors.ResponsePayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
