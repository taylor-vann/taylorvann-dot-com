package constants

import "os"

type CertAddressList struct {
	Key  		string
	Cert 		string
	Script	string
}

type PortAddressList struct {
	Http  string
	Https string
}

const (
	SessionTokenCookie 			= "X-SESSION-TOKEN"
	CsrfTokenCookie    			= "X-CSRF-TOKEN"
	DocumentTokenHeader     = "X-DOCUMENT-TOKEN"
	DocumentCsrfTokenHeader = "X-DOCUMENT-CSRF-TOKEN"

	// Ports
	HttpPort       = ":80"
	HttpsPort      = ":443"

	// Cert Locations
	certLocal       = "/usr/local/certs/gateway/https-server.crt"
	keyLocal        = "/usr/local/certs/gateway/https-server.key"
	certLetsEncrypt = "/etc/letsencrypt/"
	keyLetsEncrypt  = "/etc/letsencrypt/"

	// Cert Script Locations
	certScript 				= "/usr/local/certs/generate_self_signed_certificate.sh"
	letsEncryptScript = "/usr/local/certs/generate_letsencrypt_certificate.sh"
)

var enviornment = os.Getenv("STAGE")

func createCertConstants() *CertAddressList {
	if enviornment == "PRODUCTION" {
		return &CertAddressList{
			Key:  keyLocal,
			Cert: certLocal,
			Script:	letsEncryptScript,
		}
	}

	return &CertAddressList{
		Key:  keyLocal,
		Cert: certLocal,
		Script: certScript,
	}
}

var Ports = &PortAddressList{
	Http: HttpPort,
	Https: HttpsPort,
}

var CertAddresses = createCertConstants()
