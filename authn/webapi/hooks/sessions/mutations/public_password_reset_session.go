package mutations

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

func CreatePublicPasswordResetSession(w http.ResponseWriter, requestBody *RequestBody) {
	validRequest, errValidRequest := validateAndRemoveGuestSession(requestBody)
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
		sessions.ComposeResetPasswordSessionParams(),
	)

	if errSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)

		marshalledJSON, errMarshal := json.Marshal(&ResponsePayload{
			SessionToken: &session.SessionToken,
			CsrfToken:    &csrfAsBase64,
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
