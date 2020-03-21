package mutations

import (
	"encoding/base64"
	e "errors"
	"net/http"

	"webapi/hooks/store/errors"
	"webapi/interfaces/jwtx"
	"webapi/sessions"
	sessionsc "webapi/sessions/constants"
)

// ReadSessionAction -
type ReadSessionAction struct {
	UserID int64 `json:"user_id"`
}

// SessionParams -
type SessionParams = errors.SessionParams

// RemoveUserAction -
type RemoveUserAction = ReadSessionAction

// ResponsePayload -
type ResponsePayload = errors.ResponsePayload

// ErrorsPayload -
type ErrorsPayload = errors.Payload

// ResponseBody -
type ResponseBody = errors.ResponseBody

var (
	// CreateGuestSessionErrorMessage -
	CreateGuestSessionErrorMessage = "error creating guest session"
	// InvalidSessionProvided -
	InvalidSessionProvided = "invalid session provided"
	// UnableToMarshalSession -
	UnableToMarshalSession = "unable to marshal session"
)

func defaultErrorResponse(w http.ResponseWriter, err error) {
	errAsStr := err.Error()
	errors.BadRequest(w, &errors.Payload{
		Default: &errAsStr,
	})
}

func validateGenericSessionToken(token *string) bool {
	if token == nil {
		return false
	}

	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(token)
	if errTokenDetails != nil {
		return false
	}
	if tokenDetails.Payload.Iss != sessionsc.TaylorVannDotCom {
		return false
	}
	if tokenDetails.Payload.Exp < sessions.GetNowAsMS() {
		return false
	}
	if tokenDetails.Payload.Exp < sessions.GetNowAsMS() {
		return false
	}
	if tokenDetails.Payload.Exp-jwtx.ThreeDaysAsMS < sessions.GetNowAsMS() {
		return false
	}
	return true
}

func validateGuestSessionToken(token *string) bool {
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(token)
	if errTokenDetails != nil {
		return false
	}
	if tokenDetails.Payload.Sub != sessionsc.Guest {
		return false
	}
	if tokenDetails.Payload.Aud != sessionsc.Public {
		return false
	}
	return true
}

// UpdateSession - valid guest session and csrf are required
func UpdateSession(p *SessionParams) (*sessions.Session, error) {
	if p == nil {
		return nil, e.New("request body is nil")
	}

	if p.CsrfToken == nil {
		return nil, e.New("invalid csrf token")
	}
	csrfToken, errCsrfToken := base64.StdEncoding.DecodeString(
		*p.CsrfToken,
	)
	if errCsrfToken != nil {
		return nil, nil
	}

	// initial screen of session token
	if !validateGenericSessionToken(p.SessionToken) {
		return nil, nil
	}

	result, errEntry := sessions.Update(
		&sessions.ValidateAndRemoveParams{
		SessionToken: p.SessionToken,
		CsrfToken:    &csrfToken,
	})
	if errEntry != nil {
		return nil, errEntry
	}
	if result != nil {
		return result, nil
	}

	return nil, nil
}
