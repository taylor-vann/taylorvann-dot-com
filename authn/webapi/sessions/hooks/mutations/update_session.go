package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
)

func UpdateSession(w http.ResponseWriter, requestBody *requests.Body) {
	userSession, errUserSession := updateGenericSession(requestBody)
	if errUserSession != nil {
		errAsStr := errUserSession.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	if userSession != nil {
		marshalledJSON, errMarshal := json.Marshal(
			&responses.SessionPayload{
				SessionToken: userSession.SessionToken,
			},
		)
		if errMarshal != nil {
			errors.BadRequest(w, &responses.ErrorsPayload{
				Session: &UnableToMarshalSession,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&marshalledJSON)
		return
	}

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &responses.ErrorsPayload{
		Session: &errors.UnableToUpdateSession,
		Default: &errorAsStr,
	})
}
