package mutations

import (
	err "errors"

	"webapi/hooks/sessions/requests"
	"webapi/interfaces/jwtx"
	"webapi/sessions"
	"webapi/sessions/constants"
)

var (
	CreateGuestSessionErrorMessage = "error creating guest session"
	InvalidRequestProvided         = "invalid request provided"
	InvalidSessionProvided         = "invalid session provided"
	UnableToMarshalSession         = "unable to marshal session"
)

func validateAndRemoveSession(requestBody *requests.Body, audience string, subject string) (bool, error) {
	if requestBody.Params == nil {
		return false, err.New("request params are nil")
	}

	tokenResults := jwtx.ValidateSessionTokenByParams(&jwtx.ValidateTokenParams{
		Token:    requestBody.Params.SessionToken,
		Issuer:		constants.TaylorVannDotCom,
		Audience: audience,
		Subject:  subject,
	})

	if !tokenResults {
		return false, nil
	}

	result, errEntry := sessions.ValidateAndRemove(
		&sessions.ValidateAndRemoveParams{
			SessionToken: requestBody.Params.SessionToken,
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

func updateGenericSession(requestBody *requests.Body) (*sessions.Session, error) {
	if requestBody == nil {
		return nil, err.New("request body is nil")
	}

	if requestBody.Params == nil {
		return nil, err.New("request params are nil")
	}

	tokenResults := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token:    requestBody.Params.SessionToken,
		Issuer:		constants.TaylorVannDotCom,
	})
	if !tokenResults {
		return nil, err.New("unable to validate generic token")
	}

	session, errSession := sessions.Update(&sessions.UpdateParams{
		SessionToken: requestBody.Params.SessionToken,
	})

	if errSession != nil {
		return nil, errSession
	}
	if session != nil {
		return session, nil
	}

	return nil, nil
}
