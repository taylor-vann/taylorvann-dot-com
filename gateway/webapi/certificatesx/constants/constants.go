// brian taylor vann
// briantaylorvann dot com

package constants

import (
	"fmt"
	"os"
)

// $ openssl req -x509 -sha256 -newkey rsa:2048 -keyout certificate.key -out certificate.crt -days 1024 -nodes -subj "/C=<Country Code>/ST=<State>/L=<City>/O=<Organization>/CN=<Common Name>"
const (
	ApplicationJsonHeader = "application/json"

	OpenSSL = "openssl"
	Req = "req"
	NewOption = "-new"
	Sha256Option = "-sha256"
	NewKeyOption = "-newkey"
	RSAEncoding = "rsa:4096"
	DaysOption = "-days"
	ThreeSixtyFive = "365"
	NodesOption = "-nodes"
	X509Option = "-x509"
	SubjectOption = "-subj"
	KeyOutOption = "-keyout"
	CertOutOption = "-out"
	SubjectStatement = "/C=%s/ST=%s/L=%s/O=%s/CN=%s"
)

var (
	Environment = os.Getenv("STAGE")

	CertsHostname = os.Getenv("CERTS_HOSTNAME")
	
	// Addresses
	AuthnAddress = os.Getenv("AUTHN_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_ADDRESS")
	MailAddress = os.Getenv("MAIL_ADDRESS")
	MediaAddress = os.Getenv("MEDIA_ADDRESS")
	LogsAddress = os.Getenv("LOGS_ADDRESS")

	// certificates
	CertFilepath = os.Getenv("CERTS_CRT_FILEPATH")
	KeyFilepath = os.Getenv("CERTS_KEY_FILEPATH")
	CountryCode = os.Getenv("CERTS_COUNTRY_CODE")
	State = os.Getenv("CERTS_STATE")
	City = os.Getenv("CERTS_CITY")
	Organization = os.Getenv("CERTS_ORGANIZATION")

	SelfSignedSubject = GetSelfSignedCertificateSubjectStatement()
)

func GetStatement(statement string, compliment string) string {
	return fmt.Sprintf(statement, compliment)
}

func GetSelfSignedCertificateSubjectStatement() string {
	return fmt.Sprintf(
		SubjectStatement,
		CountryCode,
		State,
		City,
		Organization,
		CertsHostname,
	)
}