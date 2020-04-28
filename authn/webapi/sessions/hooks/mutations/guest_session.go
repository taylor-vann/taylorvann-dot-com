package mutations

import (
	"encoding/json"
	"net/http"
	
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func CreateGuestSession(w http.ResponseWriter, requestBody *requests.Body) {	
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.SessionParams
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	session, errSession := sessionsx.Create(&sessionsx.CreateParams{
		Environment: params.Environment,
		Claims: *sessionsx.CreateGuestSessionClaims(),
	})

	if errSession == nil {
		payload := responses.Session{
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
	errors.BadRequest(w, &responses.Errors{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
