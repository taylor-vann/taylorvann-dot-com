package queries

import (
	"encoding/json"
	"net/http"

	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"

	"github.com/taylor-vann/tvgtb/jwtx"
)

// requires no cookies or authorization
func ValidateGuestSession(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			RequestBody: &errors.BadRequestFail,
		})
		return
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		sessionCookie.Value,
	)
	if errTokenDetails != nil {
		errors.DefaultErrorResponse(w, errTokenDetails)
		return
	}
	if tokenDetails.Payload.Aud != "public" || tokenDetails.Payload.Sub != "guest" {
		errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.ValidateGuest
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {		
		errors.DefaultErrorResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errReadSession := sessionsx.Read(&requests.Read{
		Environment: params.Environment,
		Token: sessionCookie.Value,
	})
	if errReadSession != nil {
		errors.DefaultErrorResponse(w, errReadSession)
		return
	}

	
	if sessionIsValid {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&responses.Body{
			Session: &responses.Session{
				Token: sessionCookie.Value,
			},
		})
		return
	}

	errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
}
