package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/requests"
	"webapi/hooks/sessions/responses"
	"webapi/sessions"
	"webapi/sessions/constants"
)

func CreateCreateAccountSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody.Params == nil || requestBody.Params.AccountCredentials == nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
		return
	}

	validRequest, errValidRequest := validateAndRemoveSession(
		requestBody,
		constants.Guest,
		constants.Document,
	)
	if errValidRequest != nil {
		errAsStr := errValidRequest.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}
	if !validRequest {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
		return
	}

	sessionParams, errSessionParams := sessions.CreateAccountCreationSessionClaims(
		&sessions.CreateUserAccountClaimsParams{
			Email: requestBody.Params.AccountCredentials.Email,
		},
	)

	if errSessionParams != nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
		return
	}

	session, errSession := sessions.Create(&sessions.CreateParams{
		Environment: requestBody.Params.Environment,
		Claims: *sessionParams,
	})

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(&errors.SessionResponsePayload{
			SessionToken: session.SessionToken,
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
	errors.BadRequest(w, &responses.ErrorsPayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
