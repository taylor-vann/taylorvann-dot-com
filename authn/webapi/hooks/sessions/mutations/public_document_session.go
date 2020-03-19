package mutations

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	
	"webapi/hooks/sessions/errors"
	"webapi/sessions"
)

func CreatePublicDocumentSession(w http.ResponseWriter, requestBody *RequestBody) {
	if requestBody.Params == nil {
		errors.CustomErrorResponse(w, errors.BadBodyFail)
	}
	if requestBody.Params.Credentials == nil {
		errors.CustomErrorResponse(w, errors.CredentialsProvidedAreNil)
	}
	publicSessionParams, errPublicSessionparams := sessions.ComposePublicDocumentSessionParams(&sessions.CreatePublicJWTParams{
		Email: *requestBody.Params.Credentials.Email,
		Password: *requestBody.Params.Credentials.Password,
	})
	if errPublicSessionparams != nil {
		errors.CustomErrorResponse(w, errors.InvalidSessionCredentials)
		return
	}

	session, errSession := sessions.Create(publicSessionParams)

	if errSession == nil {
		csrfAsBase64 := base64.StdEncoding.EncodeToString(session.CsrfToken)

		payload := ResponsePayload{
			SessionToken: &session.SessionToken,
			CsrfToken:    &csrfAsBase64,
		}
		body := ResponseBody{
			Session: &payload,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&body)
		return
	}

	errorAsStr := errSession.Error()
	errors.BadRequest(w, &errors.ResponsePayload{
		Session: &CreateGuestSessionErrorMessage,
		Default: &errorAsStr,
	})
}
