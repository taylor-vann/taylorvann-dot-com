package mutations

import (
	// "encoding/base64"
	// e "errors"
	"net/http"

	"webapi/hooks/store/errors"
	// "webapi/sessions"
	// sessionsc "webapi/sessions/constants"
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

// // UpdateSession - valid guest session and csrf are required
// func UpdateSession(p *SessionParams) (*sessions.Session, error) {
// 	if p == nil {
// 		return nil, e.New("request body is nil")
// 	}

// 	// initial screen of session token
// 	if !validateGenericSessionToken(p.SessionToken) {
// 		return nil, nil
// 	}

// 	result, errEntry := sessions.Update(
// 		&sessions.ValidateAndRemoveParams{
// 		SessionToken: p.SessionToken,
// 	})
// 	if errEntry != nil {
// 		return nil, errEntry
// 	}
// 	if result != nil {
// 		return result, nil
// 	}

// 	return nil, nil
// }
