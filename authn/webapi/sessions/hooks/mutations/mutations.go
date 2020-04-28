package mutations

import (
	err "errors"

	"webapi/sessions/hooks/requests"
	"webapi/interfaces/jwtx"
	"webapi/sessions/sessionsx"
	"webapi/sessions/constants"
)

var (
	CreateGuestSessionErrorMessage = "error creating guest session"
	InvalidRequestProvided         = "invalid request provided"
	InvalidSessionProvided         = "invalid session provided"
	UnableToMarshalSession         = "unable to marshal session"
)

func updateGenericSession(p *requests.Update) (*sessionsx.Session, error) {
	if p == nil {
		return nil, err.New("request body is nil")
	}

	tokenResults := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token:    p.SessionToken,
		Issuer:		constants.TaylorVannDotCom,
	})
	if !tokenResults {
		return nil, err.New("unable to validate generic token")
	}

	session, errSession := sessionsx.Update(p)
	if errSession != nil {
		return nil, errSession
	}

	return session, nil
}
