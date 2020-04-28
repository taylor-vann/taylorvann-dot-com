package queries

import (
	"net/http"
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func ValidateSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.UnableToValidateSession,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Read)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.UnableToValidateSession,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&params)
	if errReadSession != nil {
		errors.DefaultErrorResponse(w, errReadSession)
		return
	}

	if sessionIsValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
}
