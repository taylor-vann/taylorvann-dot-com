package queries

import (
	"net/http"
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

type RequestPayload struct {
	Environment	 string `json:"environment"`
	SessionToken string `json:"session_token"`
}

type RequestBody struct {
	Action 			 string         `json:"action"`
	Params			 *RequestPayload `json:"params"`
}

func ValidateSession(w http.ResponseWriter, requestBody *RequestBody) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.SessionProvidedIsNil)
		return
	}

	sessionIsValid, errReadSession := sessions.Read(&sessions.ReadParams{
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
