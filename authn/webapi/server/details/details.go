// brian taylor vann
// briantaylorvann dot com

package details

import (
	"encoding/json"
	"net/http"

	"io/ioutil"
)

type VersionDetails struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Build int64 `json:"build"`
}

type AuthnDetails struct {
	Service string         `json:"service"`
	Build   string         `json:"build"`
	Version VersionDetails `json:"version"`
}

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"

	initFilname = "/root/go/src/webapi/server_details.json"
)

var (
	serverDetails = ReadDetailsFromFile()
)

func ReadDetailsFromFile() *AuthnDetails {
	initJSON, errInitFile := ioutil.ReadFile(initFilname)
	if errInitFile != nil {
		return nil
	}

	var details AuthnDetails
	errDetails := json.Unmarshal(initJSON, &details)
	if errDetails != nil {
		return nil
	}

	return &details
}

func Details(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, ApplicationJson)
	json.NewEncoder(w).Encode(serverDetails)
}
