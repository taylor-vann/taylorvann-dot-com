package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func CreateCreateAccountSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.InvalidSessionCredentials,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.AccountParams)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.InvalidSessionCredentials,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	sessionParams, errSessionParams := sessionsx.CreateAccountCreationSessionClaims(
		&sessionsx.AccountParams{
			Email: params.Email,
		},
	)

	if errSessionParams != nil {
		errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionParams,
	})

	if errSession == nil {
		marshalledJSON, errMarshal := json.Marshal(&responses.Session{
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
	errors.BadRequest(w, &responses.Errors{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
