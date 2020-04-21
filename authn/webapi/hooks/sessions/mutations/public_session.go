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

func CreatePublicSession(w http.ResponseWriter, requestBody *requests.Body) {
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

	userSessionToken, errUserSessionToken := sessions.CreateUserSessionClaims(
		&sessions.CreateUserClaimsParams{
			UserID: requestBody.Params.UserCredentials.UserID,
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

	userSession, errUserSession := sessions.Create(&sessions.CreateParams{
		Claims:	*userSessionToken,
	})

	if errUserSession == nil {
		marshalledJSON, errMarshal := json.Marshal(
			&responses.SessionPayload{
				SessionToken: userSession.SessionToken,
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

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &responses.ErrorsPayload{
		Session: &errors.UnableToCreatePublicSession,
		Default: &errorAsStr,
	})
}
