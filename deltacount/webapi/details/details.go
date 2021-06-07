// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"

	"webapi/redisx"
)

type ServerDetails struct {
	HTTPSPort		 int `json:"https_port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type LimiterDetails struct {
	Interval          int `json:"interval"`
	MaxServerRequests int `json:"write_timeout"`
	MaxClientRequests int `json:"max_client_requests"`
}

type CertPaths struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"private_key"`
}

type GatewayDetails struct {
	Server    ServerDetails     `json:"server"`
	CertPaths CertPaths         `json:"cert_paths"`
	Cache     redisx.Config     `json:"cache"`
	Limiter   LimiterDetails    `json:"limiter"`
	GuestList []string          `json:"guest_list"`
}

const (
	detailsPath = "/usr/local/config/details.init.json"
)

var (
	Details, DetailsErr = ReadDetailsFromFile(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseDetails(detailsJSON *[]byte, err error) (*GatewayDetails, error) {
	if err != nil {
		return nil, err
	}

	var details GatewayDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*GatewayDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
