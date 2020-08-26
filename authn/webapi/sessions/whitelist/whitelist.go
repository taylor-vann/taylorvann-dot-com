package whitelist

import (
	"encoding/json"
	"errors"

	"github.com/taylor-vann/weblog/toolbox/golang/graylistx"
	
	"webapi/sessions/whitelist/constants"
)

type MilliSeconds = int64

type Entry struct {
	Signature	 string				`json:"signature"`
	SessionKey []byte       `json:"session_key"`
	CreatedAt  MilliSeconds `json:"created_at"`
	Lifetime   MilliSeconds `json:"expires_at"`
}

type CreateEntryParams struct {
	Environment	string				`json:"environment"`
	Signature		string				`json:"signature"`
	SessionKey	[]byte				`json:"session_key"`
	CreatedAt		MilliSeconds	`json:"created_at"`
	Lifetime		MilliSeconds	`json:"lifetime"`
}

type ReadEntryParams struct {
	Environment	string `json:"environment"`
	Signature		string `json:"signature"`
}

type RemoveEntryParams = ReadEntryParams

var (
	DayAsMS = int64(1000 * 60 * 60 * 24)
	ThreeDaysAsMS = 3 * DayAsMS
	config = graylistx.Config{
		Host:        constants.Env.Host,
		Port:        constants.Env.Port,
		Protocol:    constants.Env.Protocol,
		MaxIdle:     constants.Env.MaxIdle,
		IdleTimeout: constants.Env.IdleTimeout,
		MaxActive:   constants.Env.MaxActive,
	}

	errNilParams = errors.New("nil parameters provided")
	errCreatingGraylist = errors.New("error creating graylist")
)

var graylist, errGraylist = graylistx.Create(&config)

func getEnvironmentKey(key string, environment string) string {
	if environment != "" {
		environment = constants.Development
	}

	return key + "_" + environment
}


func CreateEntry(p *CreateEntryParams) (*Entry, error) {
	if p == nil {
		return nil, errNilParams
	}
	if graylist == nil {
		return nil, errCreatingGraylist
	}

	entry := Entry{
		Signature:	p.Signature,
		SessionKey: p.SessionKey,
		CreatedAt:  p.CreatedAt,
		Lifetime:   p.Lifetime,
	}

	entryAsJSON, errEntryAsJSON := json.Marshal(entry)
	if errEntryAsJSON != nil {
		return nil, errEntryAsJSON
	}

	environmentKey := getEnvironmentKey(p.Signature, p.Environment)
	whitelistResult, errWhitelist := graylist.SetAndExpire(
		&graylistx.SetAndExpireParams{
			Key: environmentKey,
			Value: entryAsJSON,
			ExpiryInMS: p.Lifetime,
		},
	)

	if whitelistResult == true {
		return &entry, errWhitelist
	}

	return nil, errWhitelist
}

func ReadEntry(p *ReadEntryParams) (*Entry, error) {
	if p == nil {
		return nil, errNilParams
	}
	if graylist == nil {
		return nil, errCreatingGraylist
	}

	environmentKey := getEnvironmentKey(p.Signature, p.Environment)
	entryAsByte, errEntryAsByte := graylist.Get(&graylistx.GetParams{
		Key: environmentKey,
	})
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

func RemoveEntry(p *RemoveEntryParams) (bool, error) {
	if p == nil {
		return false, errNilParams
	}
	if graylist == nil {
		return false, errCreatingGraylist
	}

	environmentKey := getEnvironmentKey(p.Signature, p.Environment)
	return graylist.Remove(&graylistx.RemoveParams{
		Key: environmentKey,
	})
}
