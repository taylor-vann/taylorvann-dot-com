package mutations

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

func CreatePublicSession(w http.ResponseWriter, requestBody *RequestBody) {
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

	userSessionToken, errUserSessionToken := sessions.ComposePublicSessionParams(
		&sessions.CreatePublicJWTParams{
			Email:    *requestBody.Params.Credentials.Email,
			Password: *requestBody.Params.Credentials.Password,
		},
	)
	if errUserSessionToken != nil {
		errorAsStr := errUserSessionToken.Error()
		errors.BadRequest(w, &errors.ResponsePayload{
			Session: &errors.InvalidSessionCredentials,
			Default: &errorAsStr,
		})
		return
	}

	userSession, errUserSession := sessions.Create(
		userSessionToken,
	)

	if errUserSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(userSession.CsrfToken)
		marshalledJSON, errMarshal := json.Marshal(&ResponsePayload{
			SessionToken: &userSession.SessionToken,
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

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &errors.ResponsePayload{
		Session: &errors.UnableToCreatePublicSession,
		Default: &errorAsStr,
	})
}
