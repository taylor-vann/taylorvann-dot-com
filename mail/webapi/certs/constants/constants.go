// brian taylor vann
// taylorvann dot com

package constants

import "os"

type CertAddressList struct {
	Cert 		string
	Key  		string
	Script	string
}

const (
	certLocal       = "/usr/local/certs/mail/https-server.crt"
	certLetsEncrypt = "/etc/letsencrypt/live/mail.taylorvann.com/fullchain.pem"

	keyLocal        = "/usr/local/certs/mail/https-server.key"
	keyLetsEncrypt  = "/etc/letsencrypt/live/mail.taylorvann.com/privkey.pem"

	scriptLocal 			= "/usr/local/certs/generate_self_signed_certificate.sh"
	scriptLetsEncrypt = "/usr/local/certs/generate_letsencrypt_certificate.sh"
)

var environment = os.Getenv("STAGE")
var Addresses = createCertConstants()

func createCertConstants() *CertAddressList {
	if environment == "PRODUCTION" {
		return &CertAddressList{
			Cert: certLetsEncrypt,
			Key:  keyLetsEncrypt,
			Script:	scriptLetsEncrypt,
		}
	}

	return &CertAddressList{
		Cert: certLocal,
		Key:  keyLocal,
		Script:	scriptLocal,
	}
}

