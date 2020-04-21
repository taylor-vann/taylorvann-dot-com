package mutations

import (
	"encoding/json"
	"net/http"
	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/responses"
)

func UpdateSession(w http.ResponseWriter, requestBody *responses.RequestBody) {
	userSession, errUserSession := updateGenericSession(requestBody)
	if errUserSession != nil {
		errAsStr := errUserSession.Error()
		errors.BadRequest(w, &responses.ErrorsResponsePayload{
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
			errors.BadRequest(w, &responses.ErrorsResponsePayload{
				Session: &UnableToMarshalSession,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&marshalledJSON)
		return
	}

	errorAsStr := errUserSession.Error()
	errors.BadRequest(w, &responses.ErrorsResponsePayload{
		Session: &errors.UnableToUpdateSession,
		Default: &errorAsStr,
	})
}
