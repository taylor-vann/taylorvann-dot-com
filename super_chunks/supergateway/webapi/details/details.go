// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"
)

type ServerDetails struct {
	HTTPPort     int `json:"http_port"`
	HTTPSPort    int `json:"https_port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type CertPaths struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"private_key"`
}

type GatewayDetails struct {
	ServiceName	string						`json:"service_name"`
	CertPaths 	CertPaths         `json:"cert_paths"`
	Routes    	map[string]string `json:"routes"`
	Server   		ServerDetails     `json:"server"`
}

const (
	detailsPath = "/usr/local/config/config.json"
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
