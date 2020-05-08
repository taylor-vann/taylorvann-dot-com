// brian taylor vann
// taylorvann dot com

package constants

import "os"

type PortAddressList struct {
	Smtp  string
	Https string
}

const (
	SmtpPort       = ":25"
	HttpsPort 		 = ":6000"
)

var Enviornment = os.Getenv("STAGE")

var Ports = &PortAddressList{
	Smtp: SmtpPort,
	Https: HttpsPort,
}
