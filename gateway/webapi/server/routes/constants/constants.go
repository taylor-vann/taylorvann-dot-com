//	brian taylor vann
//	briantaylorvann dot com

package constants

import (
	"os"
)

type SubDomain = string
type Address = string
type DomainDetailsMap = map[SubDomain]Address

const (
	BrianTaylorVann = "briantaylorvann"
	Authn	= "authn"
	Logs 	= "logs"
	Mail	= "mail"
	Media = "media"
)

var (
	Enviornment = os.Getenv("STAGE")

	CertsHostname = os.Getenv("CERTS_HOSTNAME")
	AuthnAddress = os.Getenv("AUTHN_URL_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_URL_ADDRESS")
	LogsAddress = os.Getenv("LOGS_URL_ADDRESS")
	MailAddress = os.Getenv("MAIL_URL_ADDRESS")
	MediaAddress = os.Getenv("MEDIA_URL_ADDRESS")
)

var Routes = createDomainDetailsMap()

func createDomainDetailsMap() *DomainDetailsMap {	
	domains := make(DomainDetailsMap)

	domains[BrianTaylorVann] = ClientsAddress

	domains[Authn] = AuthnAddress
	domains[Logs] = LogsAddress
	domains[Mail] = MailAddress
	domains[Media] = MediaAddress

	return &domains
}