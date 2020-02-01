// Package hooks - ping
//
// brian taylor vann
// taylorvann dot com
package hooks

import (
	"encoding/json"
	"net/http"
)

type versionDetails struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Build int64 `json:"build"`
}

type authenticationDetails struct {
	Service string         `json:"service"`
	Version versionDetails `json:"version"`
}

var version = versionDetails{
	Major: 0,
	Minor: 1,
	Build: 1,
}

const service = "taylorvann_authentication"

var authnDetails = authenticationDetails{
	Service: service,
	Version: version,
}

// Ping - get information about our api
func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authnDetails)
}
