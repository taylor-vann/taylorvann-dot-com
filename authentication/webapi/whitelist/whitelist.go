package whitelist

import (
	"encoding/json"
	"errors"
	"webapi/interfaces/whitelistx"
	"webapi/utils"
)

// MilliSecond -
type MilliSecond = int64

// Entry -
type Entry struct {
	CsrfToken  []byte             `json:"csrf_token"`
	SessionKey []byte             `json:"session_key"`
	CreatedAt  utils.MilliSeconds `json:"created_at"`
	Lifetime   utils.MilliSeconds `json:"expires_at"`
}

// CreateEntryParams -
type CreateEntryParams struct {
	CreatedAt  utils.MilliSeconds
	Lifetime   utils.MilliSeconds
	SessionKey *[]byte
	Signature  *string
}

// ReadEntryParams -
type ReadEntryParams struct {
	Signature *string
}

// DayAsMS -
var DayAsMS = int64(1000 * 60 * 60 * 24)

// ThreeDaysAsMS -
var ThreeDaysAsMS = 3 * DayAsMS

// RemoveEntryParams -
type RemoveEntryParams = ReadEntryParams

// CreateEntry -
func CreateEntry(p *CreateEntryParams) (*Entry, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	csrfToken, errCsrfToken := utils.GenerateRandomByteArray(128)
	if errCsrfToken != nil {
		return nil, errCsrfToken
	}

	entry := Entry{
		CsrfToken:  *csrfToken,
		SessionKey: *(p.SessionKey),
		CreatedAt:  p.CreatedAt,
		Lifetime:   p.Lifetime,
	}

	// marshal entry to byte array
	entryAsJSON, errEntryAsJSON := json.Marshal(entry)
	if errEntryAsJSON != nil {
		return nil, errEntryAsJSON
	}

	// save to whitelist
	whitelistResult, errWhitelist := whitelistx.SetAndExpire(
		*(p.Signature),
		&entryAsJSON,
		p.Lifetime,
	)

	if errWhitelist != nil {
		return nil, errWhitelist
	}

	if whitelistResult == true {
		return &entry, errWhitelist
	}

	return nil, errCsrfToken
}

// ReadEntry -
func ReadEntry(p *ReadEntryParams) (*Entry, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	entryAsByte, errEntryAsByte := whitelistx.Get(*(p.Signature))
	if errEntryAsByte != nil {
		return nil, errEntryAsByte
	}

	var entry Entry
	errUnmarshal := json.Unmarshal(*entryAsByte, &entry)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &entry, errUnmarshal
}

// RemoveEntry -
func RemoveEntry(p *RemoveEntryParams) (bool, error) {
	if p == nil {
		return false, errors.New("nil parameters provided")
	}
	return whitelistx.Remove(*(p.Signature))
}
