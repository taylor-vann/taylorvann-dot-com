package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
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

	sessionParams, errSessionParams := sessionsx.CreateAccountCreationSessionClaims(
		&sessionsx.CreateUserAccountClaimsParams{
			Email: requestBody.Params.AccountCredentials.Email,
		},
	)

	if errSessionParams != nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: requestBody.Params.Environment,
		Claims: *sessionParams,
	})

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(&responses.SessionPayload{
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
