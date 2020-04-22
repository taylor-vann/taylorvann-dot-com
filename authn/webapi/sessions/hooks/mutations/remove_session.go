package mutations

import (
	"net/http"

	"webapi/interfaces/jwtx"
	"webapi/sessions/hooks/errors"
	"webapi/sessions/hooks/requests"
	"webapi/sessions/hooks/responses"
	"webapi/sessions/sessionsx"
)

func RemoveSession(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
	}

	token, errSignature := jwtx.RetrieveTokenFromString(
		requestBody.Params.SessionToken,
	)
	if errSignature != nil {
		errAsStr := errSignature.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	result, errResponseBody := sessionsx.Remove(
		&sessionsx.RemoveParams{
			Signature: token.Signature,
		},
	)

	if errResponseBody != nil {
		errAsStr := errResponseBody.Error()
		errors.BadRequest(w, &responses.ErrorsPayload{
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
