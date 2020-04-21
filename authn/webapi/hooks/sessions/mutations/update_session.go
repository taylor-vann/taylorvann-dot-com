package mutations

import (
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
		marshalledJSON, errMarshal := json.Marshal(
			&errors.SessionResponsePayload{
				SessionToken: userSession.SessionToken,
			},
		)
		if errMarshal != nil {
			errors.BadRequest(w, &errors.ResponsePayload{
				Session: &UnableToMarshalSession,
			})
			return
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
