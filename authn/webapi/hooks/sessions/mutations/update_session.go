package mutations

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"webapi/hooks/sessions/errors"
)

func UpdateSession(w http.ResponseWriter, requestBody *RequestBody) {
	userSession, errUserSession := updateGenericSession(requestBody)
	if errUserSession != nil {
		errAsStr := errUserSession.Error()
		errors.BadRequest(w, &errors.ResponsePayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	if userSession != nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(userSession.CsrfToken)
		marshalledJSON, errMarshal := json.Marshal(&ResponsePayload{
			SessionToken: &userSession.SessionToken,
			CsrfToken:    &csrfAsBase64,
		})
		if errMarshal != nil {
			errors.BadRequest(w, &errors.ResponsePayload{
				Session: &UnableToMarshalSession,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&marshalledJSON)
		return
	}

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &errors.ResponsePayload{
		Session: &errors.UnableToUpdateSession,
		Default: &errorAsStr,
	})
}
