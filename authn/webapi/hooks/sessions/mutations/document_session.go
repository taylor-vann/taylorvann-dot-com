package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/responses"
	"webapi/sessions"
)

func CreateDocumentSession(w http.ResponseWriter) {
	session, errSession := sessions.Create(&sessions.CreateParams{
		Claims: *sessions.CreateDocumentSessionClaims(),
	})

	if errSession == nil {
		payload := responses.SessionPayload{
			SessionToken: session.SessionToken,
		}
		body := responses.Body{
			Session: &payload,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&body)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &responses.ErrorsPayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
