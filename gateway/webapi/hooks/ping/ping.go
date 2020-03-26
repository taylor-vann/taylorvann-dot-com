// brian taylor vann
// taylorvann dot com

// Package - ping
package ping

import (
	"encoding/json"
	"net/http"
)

// VersionDetails -
type VersionDetails struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Build int64 `json:"build"`
}

// GatewayDetails -
type GatewayDetails struct {
	Service string         `json:"service"`
	Build   string         `json:"build"`
	Version VersionDetails `json:"version"`
}

var version = VersionDetails{
	Major: 0,
	Minor: 1,
	Build: 1,
}

const service = "taylorvann_gateway"
const build = "single_server"

var authnDetails = GatewayDetails{
	Service: service,
	Build:   build,
	Version: version,
}

// Details - get information about our api
func Details(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authnDetails)
}
