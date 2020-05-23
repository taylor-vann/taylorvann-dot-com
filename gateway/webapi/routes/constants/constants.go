//	brian taylor vann
//	taylorvann dot com

package constants

import (
	"os"
)

type SubDomain = string
type Address = string
type DomainDetailsMap = map[SubDomain]Address

const (
	Briantaylorvann = "briantaylorvann"
	Authn	= "authn"
	Logs 	= "logs"
	Mail	= "mail"
	media = "media"
)

var (
	Enviornment = os.Getenv("STAGE")

	CertsHostname = os.Getevn("CERTS_HOSTNAME")
	AuthnAddress = os.Getenv("AUTHN_URL_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_URL_ADDRESS")
	LogsAddress = os.Getenv("LOGS_URL_ADDRESS")
	MailAddress = os.Getenv("MAIL_URL_ADDRESS")
	MediaAddress = os.Getenv("MEDIA_URL_ADDRESS")
)

var Routes = createDomainDetailsMap()

func createDomainDetailsMap() *DomainDetailsMap {	
	domains := make(DomainDetailsMap)

	domains[briantaylorvann] = ClientsAddress

	domains[authn] = AuthnAddress
	domains[logs] = LogsAddress
	domains[mail] = MailAddress
	domains[media] = MediaAddress

	return &domains
}