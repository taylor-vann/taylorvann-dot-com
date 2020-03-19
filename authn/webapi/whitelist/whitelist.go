package whitelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"webapi/interfaces/whitelistx"
)

// MilliSeconds -
type MilliSeconds = int64

// Entry -
type Entry struct {
	CsrfToken  []byte       `json:"csrf_token"`
	SessionKey []byte       `json:"session_key"`
	CreatedAt  MilliSeconds `json:"created_at"`
	Lifetime   MilliSeconds `json:"expires_at"`
}

// CreateEntryParams -
type CreateEntryParams struct {
	SessionKey *[]byte
	Signature  *string
	CreatedAt  MilliSeconds
	Lifetime   MilliSeconds
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

// generateRandomByteArray -
func generateRandomByteArray(n uint32) (*[]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// CreateEntry -
func CreateEntry(p *CreateEntryParams) (*Entry, error) {
	if p == nil {
		return nil, errors.New("nil parameters provided")
	}

	csrfToken, errCsrfToken := generateRandomByteArray(128)
	if errCsrfToken != nil {
		return nil, errCsrfToken
	}

	entry := Entry{
		CsrfToken:  *csrfToken,
		SessionKey: *p.SessionKey,
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
		fmt.Println("err whitelist")
		fmt.Println(errWhitelist)
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

	if entryAsByte == nil {
		return nil, nil
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
