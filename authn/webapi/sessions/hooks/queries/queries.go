package queries

import (
	"encoding/json"
	"net/http"

	"github.com/taylor-vann/tvgtb/jwtx"
	
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

// requires no cookies or authorization
func ValidateGuestSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&params)
	if errReadSession != nil {
		errors.DefaultErrorResponse(w, errReadSession)
		return
	}

	if sessionIsValid == false {
		errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
		return
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		params.Token,
	)
	if errTokenDetails != nil {
		errors.DefaultErrorResponse(w, errTokenDetails)
		return
	}

	if tokenDetails.Payload.Aud == "public" && tokenDetails.Payload.Sub == "guest" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&responses.Body{
			Session: &responses.Session{
				Token: params.Token,
			},
		})
		return
	}

	errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
}
