package mutations

import (
	"encoding/base64"
	"errors"

	sessionErrors "webapi/hooks/sessions/errors"
	"webapi/interfaces/jwtx"
	"webapi/sessions"
	"webapi/sessions/constants"
)

type RequestPayload = sessionErrors.RequestPayload
type RequestBody = sessionErrors.RequestBody
type ResponsePayload = sessionErrors.SessionResponsePayload
type ErrorsPayload = sessionErrors.ResponsePayload
type ResponseBody = sessionErrors.ResponseBody

var (
	CreateGuestSessionErrorMessage = "error creating guest session"
	InvalidSessionProvided         = "invalid session provided"
	UnableToMarshalSession         = "unable to marshal session"
)

func validateGenericSessionToken(token *string) bool {
	nowAsMS := sessions.GetNowAsMS()
	threeDaysAgoAsMS := nowAsMS - jwtx.ThreeDaysAsMS
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(token)
	if errTokenDetails == nil &&
		tokenDetails.Payload.Iss == constants.TaylorVannDotCom &&
		threeDaysAgoAsMS < tokenDetails.Payload.Exp &&
		tokenDetails.Payload.Iat < nowAsMS {
		return true
	}

	return false
}

func validateGuestSessionToken(token *string) bool {
	if !validateGenericSessionToken(token) {
		return false
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(token)
	if errTokenDetails == nil && tokenDetails.Payload.Sub == constants.Guest {
		return true
	}

	return false
}

func validateAndRemoveGuestSession(requestBody *RequestBody) (bool, error) {
	if requestBody.Params == nil {
		return false, errors.New("request params are nil")
	}

	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(
		*requestBody.Params.CsrfToken,
	)
	if errCsrfToken != nil {
		return false, nil
	}

	if !validateGuestSessionToken(requestBody.Params.SessionToken) {
		return false, nil
	}

	result, errEntry := sessions.ValidateAndRemove(
		&sessions.ValidateAndRemoveParams{
			SessionToken: requestBody.Params.SessionToken,
			CsrfToken:    &csrfToken,
		},
	)
	if errEntry != nil {
		return false, errEntry
	}
	if result != nil {
		return true, nil
	}

	return false, nil
}

func validateGenericSession(requestBody *RequestBody) (bool, error) {
	if requestBody.Params == nil {
		return false, errors.New("request params are nil")
	}

	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(
		*requestBody.Params.CsrfToken,
	)
	if errCsrfToken != nil {
		return false, nil
	}

	if !validateGenericSessionToken(requestBody.Params.SessionToken) {
		return false, nil
	}

	result, errEntry := sessions.ValidateAndRemove(&sessions.ValidateAndRemoveParams{
		SessionToken: requestBody.Params.SessionToken,
		CsrfToken:    &csrfToken,
	})
	if errEntry != nil {
		return false, errEntry
	}
	if result != nil {
		return true, nil
	}

	return false, nil
}

func updateGenericSession(requestBody *RequestBody) (*sessions.Session, error) {
	if requestBody == nil {
		return nil, errors.New("request body is nil")
	}

	if requestBody.Params == nil {
		return nil, errors.New("request params are nil")
	}

	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(
		*requestBody.Params.CsrfToken,
	)
	if errCsrfToken != nil {
		return nil, errCsrfToken
	}

	if !validateGenericSessionToken(requestBody.Params.SessionToken) {
		return nil, errors.New("error validating session token")
	}

	session, errSession := sessions.Update(&sessions.UpdateParams{
		SessionToken: requestBody.Params.SessionToken,
		CsrfToken:    &csrfToken,
	})

	if errSession != nil {
		return nil, errSession
	}
	if session != nil {
		return session, nil
	}

	return nil, nil
}
