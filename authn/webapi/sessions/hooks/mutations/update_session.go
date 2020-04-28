package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
)

func UpdateSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.UnableToUpdateSession,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Update
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	userSession, errUserSession := updateGenericSession(&params)
	if errUserSession != nil {
		errAsStr := errUserSession.Error()
		errors.BadRequest(w, &responses.Errors{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	if userSession != nil {
		marshalledJSON, errMarshal := json.Marshal(
			&responses.Session{
				SessionToken: userSession.SessionToken,
			},
		)
		if errMarshal != nil {
			errors.BadRequest(w, &responses.Errors{
				Session: &UnableToMarshalSession,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&marshalledJSON)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Session: &errors.UnableToUpdateSession,
	})
}
