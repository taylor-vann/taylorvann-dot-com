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
	briantaylorvann = "briantaylorvann"
	authn		= "authn"
	logs 		= "logs"
	mail		= "mail"
	statics = "statics"
)

var (
	Enviornment = os.Getenv("STAGE")

	AuthnAddress = os.Getenv("AUTHN_URL_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_URL_ADDRESS")
	LogsAddress = os.Getenv("LOGS_URL_ADDRESS")
	MailAddress = os.Getenv("MAIL_URL_ADDRESS")
	StaticAddress = os.Getenv("STATIC_URL_ADDRESS")
)

var Routes = createDomainDetailsMap()

func createDomainDetailsMap() *DomainDetailsMap {	
	domains := make(DomainDetailsMap)

	domains[briantaylorvann] = ClientsAddress

	domains[authn] = AuthnAddress
	domains[logs] = LogsAddress
	domains[mail] = MailAddress
	domains[statics] = StaticAddress

	return &domains
}