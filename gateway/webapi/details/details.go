// brian taylor vann
// briantaylorvann dot com

package details

import (
	"encoding/json"
	"io/ioutil"
)

type CertPaths struct {
	Cert				string	`json:"cert"`
	PrivateKey	string	`json:"private_key"`
}

type GatewayDetails struct {
	CertPaths	*CertPaths				`json:"cert_paths,omitempty"`
	Routes		map[string]string	`json:"routes"`
}

const (
	detailsPath = "/usr/local/config/details.init.json"
)

var (
	Details, DetailsErr = readDetailsFromFile()
)

func readFile(path string) (*[]byte, *error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(detailsPath)
	return &detailsJSON, &errDetiailsJSON
}

func rarseDetails(detailsJSON *[]byte, err *error) (*GatewayDetails, *error) {
	if (*err != nil) {
		return nil, err;
	}

	var details GatewayDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, &errDetails
}

func readDetailsFromFile() (*GatewayDetails, *error) {
	detailsJSON, errDetailsJSON := readFile(detailsPath)
	return parseDetails(detailsJSON, errDetailsJSON)
}
