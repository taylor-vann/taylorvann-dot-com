// brian taylor vann
// taylorvann dot com
//
// Package sessions - meta-interface for whitelist, csrfs, and session tokens

package sessions

import (
	"errors"
	"fmt"
	"time"
	"webapi/interfaces/jwtx"
	"webapi/sessions/constants"
	"webapi/store"
	"webapi/whitelist"
)

// MilliSeconds -
type MilliSeconds = int64

// Session -
type Session struct {
	SessionToken string
	CsrfToken    []byte
}

// CreatePublicJWTParams -
type CreatePublicJWTParams struct {
	Email    *string
	Password *string
}

// CreateSessionParams -
type CreateSessionParams struct {
	Issuer   string
	Subject  string
	Audience string
}

// UpdateSessionParams -
type UpdateSessionParams struct {
	SessionToken *string
	CsrfToken    *[]byte
}

// RemoveSessionParams -
type RemoveSessionParams = whitelist.RemoveEntryParams

// getNowAsMS -
func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func composeJWTClaims(p *CreateSessionParams, issuedAt int64, expiresAt int64) *jwtx.Claims {
	jwtClaims := jwtx.Claims{
		Iss: p.Issuer,
		Sub: p.Subject,
		Aud: p.Audience,
		Iat: issuedAt,
		Exp: expiresAt,
	}

	return &jwtClaims
}

func validateCsrfTokens(external *[]byte, whitelist *[]byte) bool {
	if len(*external) == len(*whitelist) {
		for index, bit := range *whitelist {
			if (*external)[index] != bit {
				return false
			}
		}

		return true
	}

	return false
}

// ComposeCreateGuestSessionParams -
func ComposeCreateGuestSessionParams() *CreateSessionParams {
	params := CreateSessionParams{
		Issuer:   constants.TaylorVannDotCom,
		Subject:  constants.Guest,
		Audience: constants.Public,
	}

	return &params
}

// ComposeCreatePublicSessionParams -
func ComposeCreatePublicSessionParams(p *CreatePublicJWTParams) (*CreateSessionParams, error) {
	// validate user through store
	userRow, errValidUser := store.ValidateUser(
		&store.ValidateUserParams{
			Email:    *(p.Email),
			Password: *(p.Password),
		},
	)
	if userRow == nil {
		return nil, errors.New("bad credentials provided")
	}
	if errValidUser != nil {
		return nil, errValidUser
	}

	params := CreateSessionParams{
		Issuer:   constants.TaylorVannDotCom,
		Subject:  fmt.Sprintf("%d", userRow.ID),
		Audience: constants.Public,
	}

	return &params, nil
}

// CreateSession -
func CreateSession(p *CreateSessionParams) (*Session, error) {
	if p == nil {
		return nil, errors.New("nil CreateSessionParams provided")
	}
	// create guest jwt
	issuedAt := getNowAsMS()
	lifetime := constants.ThreeDaysAsMS
	expiresAt := issuedAt + lifetime

	token, errToken := jwtx.CreateJWT(
		composeJWTClaims(p, issuedAt, expiresAt),
	)
	if errToken != nil {
		return nil, errToken
	}

	// add entry
	entry, errEntry := whitelist.CreateEntry(
		&whitelist.CreateEntryParams{
			CreatedAt:  issuedAt,
			Lifetime:   lifetime,
			SessionKey: token.RandomSecret,
			Signature:  &token.Token.Signature,
		},
	)
	if errEntry != nil {
		return nil, errEntry
	}

	// compose session
	sessionTokenAsStr, errSessionTokenAsStr := jwtx.ConvertTokenToString(
		token.Token,
	)
	if errSessionTokenAsStr != nil {
		return nil, errSessionTokenAsStr
	}

	session := Session{
		SessionToken: *sessionTokenAsStr,
		CsrfToken:    entry.CsrfToken,
	}

	return &session, nil
}

// UpdateSessionIfExists -
func UpdateSessionIfExists(p *UpdateSessionParams) (*Session, error) {
	sessionDetails, errSessionDetails := jwtx.RetrieveTokenDetailsFromString(
		p.SessionToken,
	)
	if errSessionDetails != nil {
		return nil, errSessionDetails
	}

	entry, errEntry := whitelist.ReadEntry(
		&whitelist.ReadEntryParams{
			Signature: sessionDetails.Signature,
		},
	)
	if entry == nil {
		return nil, nil
	}
	if errEntry != nil {
		return nil, errEntry
	}

	isValidCsrfToken := validateCsrfTokens(p.CsrfToken, &entry.CsrfToken)
	if isValidCsrfToken {
		removeResult, errRemoveResult := whitelist.RemoveEntry(
			&whitelist.RemoveEntryParams{
				Signature: sessionDetails.Signature,
			},
		)
		if errRemoveResult != nil {
			return nil, errRemoveResult
		}
		if removeResult == false {
			return nil, errRemoveResult
		}

		// validated, removed, return new token
		return CreateSession(&CreateSessionParams{
			Issuer:   sessionDetails.Payload.Iss,
			Subject:  sessionDetails.Payload.Sub,
			Audience: sessionDetails.Payload.Aud,
		})
	}

	return nil, nil
}

// RemoveSession -
func RemoveSession(p *RemoveSessionParams) (bool, error) {
	return whitelist.RemoveEntry(p)
}
