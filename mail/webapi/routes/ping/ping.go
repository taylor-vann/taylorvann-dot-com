// brian taylor vann
// taylorvann dot com

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

type MailDetails struct {
	Service string         `json:"service"`
	Build   string         `json:"build"`
	Version VersionDetails `json:"version"`
}

var version = VersionDetails{
	Major: 0,
	Minor: 1,
	Build: 1,
}

const service = "mail"
const build = "single_server"

var authnDetails = MailDetails{
	Service: service,
	Build:   build,
	Version: version,
}

func Details(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authnDetails)
}
