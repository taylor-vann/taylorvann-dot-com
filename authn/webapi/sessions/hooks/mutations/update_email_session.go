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

func CreateUpdateEmailSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody.Params == nil || requestBody.Params.AccountCredentials == nil {
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

	userSessionToken, errUserSessionToken := sessionsx.CreateUpdatePasswordSessionClaims(
		&sessionsx.CreateUserAccountClaimsParams{
			Email: requestBody.Params.AccountCredentials.Email,
		},
	)
	if errUserSessionToken != nil {
		errorAsStr := errUserSessionToken.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &errors.InvalidSessionCredentials,
			Default: &errorAsStr,
		})
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Claims: *userSessionToken,
	})

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(
			&responses.SessionPayload{
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
	errors.BadRequest(w, &responses.ErrorsPayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
