package mutations

import (
	"net/http"
	"webapi/sessions"

	"webapi/hooks/sessions/errors"
	"webapi/hooks/sessions/responses"

	"webapi/interfaces/jwtx"
)

func RemoveSession(w http.ResponseWriter, requestBody *responses.RequestBody) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, InvalidSessionProvided)
	}

	token, errSignature := jwtx.RetrieveTokenFromString(
		requestBody.Params.SessionToken,
	)
	if errSignature != nil {
		errAsStr := errSignature.Error()
		errors.BadRequest(w, &responses.ErrorsResponsePayload{
			Session: &InvalidSessionProvided,
			Default: &errAsStr,
		})
		return
	}

	result, errResponseBody := sessions.Remove(
		&sessions.RemoveParams{
			Signature: token.Signature,
		},
	)

	if errResponseBody != nil {
		errAsStr := errResponseBody.Error()
		errors.BadRequest(w, &responses.ErrorsResponsePayload{
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
