// brian taylor vann
// briantaylorvann dot com

package ping

import (
	"encoding/json"
	"net/http"
)

type VersionDetails struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Build int64 `json:"build"`
}

type AuthenticationDetails struct {
	Service string         `json:"service"`
	Build   string         `json:"build"`
	Version VersionDetails `json:"version"`
}

var version = VersionDetails{
	Major: 0,
	Minor: 1,
	Build: 2,
}

const service = "authn"
const build = "single_server"

var authnDetails = AuthenticationDetails{
	Service: service,
	Build:   build,
	Version: version,
}

func Details(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authnDetails)
}
