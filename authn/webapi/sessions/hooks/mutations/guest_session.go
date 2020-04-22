package mutations

import (
	"encoding/json"
	"net/http"
	
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func CreateGuestSession(w http.ResponseWriter) {	
	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Claims: *sessionsx.CreateGuestSessionClaims(),
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
