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

type ServiceDetails struct {
	Service string         `json:"service"`
	Build   string         `json:"build"`
	Version VersionDetails `json:"version"`
}

var version = VersionDetails{
	Major: 0,
	Minor: 1,
	Build: 1,
}

const service = "clients"
const build = "single_server"

var clientsDetails = ServiceDetails{
	Service: service,
	Build:   build,
	Version: version,
}

func Details(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clientsDetails)
}