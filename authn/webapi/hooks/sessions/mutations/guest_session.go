package mutations

import (
	"encoding/json"
	"net/http"
	
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

func CreateGuestSession(w http.ResponseWriter) {	
	session, errSession := sessions.Create(
		sessions.ComposeGuestSessionParams(),
	)

	if errSession == nil {
		payload := ResponsePayload{
			SessionToken: &session.SessionToken,
		}
		body := ResponseBody{
			Session: &payload,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&body)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &errors.ResponsePayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}