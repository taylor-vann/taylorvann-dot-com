package mutations

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

func CreateGuestDocumentSession(w http.ResponseWriter) {	
	session, errSession := sessions.Create(
		sessions.ComposeGuestDocumentSessionParams(),
	)

	if errSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)

		payload := ResponsePayload{
			SessionToken: &session.SessionToken,
			CsrfToken:    &csrfAsBase64,
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
