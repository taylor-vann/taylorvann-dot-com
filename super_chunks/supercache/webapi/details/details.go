// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"

	"webapi/redisx"
)

type ServerDetails struct {
	HTTPPort		 int `json:"http_port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type SuperCacheDetails struct {
	ServiceName string				`json:"service_name"`	
	Server    	ServerDetails	`json:"server"`
	Cache     	redisx.Config	`json:"cache"`
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

func parseDetails(detailsJSON *[]byte, err error) (*SuperCacheDetails, error) {
	if err != nil {
		return nil, err
	}

	var details SuperCacheDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*SuperCacheDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
