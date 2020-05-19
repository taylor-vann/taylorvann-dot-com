//	brian taylor vann
//	taylorvann dot com

package constants

import (
	"os"

	"github.com/taylor-vann/tvgtb/certificatesx/constants"
)

const (
	Dns 	= ":53"
	Http  = ":80"
	Https = ":443"
)

var (
	Environment = os.Getenv("ENVIRONMENT")

	CertFilepath = getCertPath()
	KeyFilepath = getKeyPath()
)

func getCertPath() string {
	if Environment == "PRODUCTION" {
		return constants.ProductionCertLetsEncryptFilepath
	}

	return constants.SandboxCertLetsEncryptFilepath
}

func getKeyPath() string {
	if Environment == "PRODUCTION" {
		return constants.ProductionKeyLetsEncryptFilepath
	}

	return constants.SandboxKeyLetsEncryptFilepath
}