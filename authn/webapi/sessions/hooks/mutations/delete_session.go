package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func DeleteSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Session: &errors.InvalidSessionCredentials,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Delete
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	result, errResponseBody := sessionsx.Delete(&params)
	if errResponseBody != nil {
		errAsStr := errResponseBody.Error()
		errors.BadRequest(w, &responses.Errors{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	if result == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	errors.CustomErrorResponse(w, InvalidSessionProvided)
}
