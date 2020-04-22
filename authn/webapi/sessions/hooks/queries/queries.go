package queries

import (
	"net/http"
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/sessionsx"
)

func ValidateSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.SessionProvidedIsNil)
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&sessionsx.ReadParams{
		Environment:  requestBody.Params.Environment,
		SessionToken: requestBody.Params.SessionToken,
	})

	if errReadSession != nil {
		errors.DefaultErrorResponse(w, errReadSession)
		return
	}

	if sessionIsValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
}
