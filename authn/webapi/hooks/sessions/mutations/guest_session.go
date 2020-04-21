package mutations

import (
	"encoding/json"
	"net/http"
	
	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/responses"
	"webapi/sessions"
)

func CreateGuestSession(w http.ResponseWriter) {	
	session, errSession := sessions.Create(&sessions.CreateParams{
		Claims: *sessions.CreateGuestSessionClaims(),
	})

	if errSession == nil {
		payload := responses.SessionResponsePayload{
			SessionToken: session.SessionToken,
		}
		body := responses.ResponseBody{
			Session: &payload,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&body)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &responses.ErrorsResponsePayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
