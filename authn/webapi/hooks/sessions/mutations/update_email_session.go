package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/sessions"
	"webapi/sessions/constants"
)

func CreateUpdateEmailSession(w http.ResponseWriter, requestBody *RequestBody) {
	if requestBody.Params == nil || requestBody.Params.Credentials == nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
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

	userSessionToken, errUserSessionToken := sessions.CreateUpdatePasswordSessionClaims(
		&sessions.CreateUserAccountClaimsParams{
			Email: requestBody.Params.Credentials.Email,
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

	session, errSession := sessions.Create(&sessions.CreateParams{
		Claims: *userSessionToken,
	})

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(
			&errors.SessionResponsePayload{
				SessionToken: session.SessionToken,
			},
		)
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
