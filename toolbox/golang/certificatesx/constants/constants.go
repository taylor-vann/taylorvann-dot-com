// brian taylor vann
// toolbox-go

package constants

import (
	"fmt"
	"os"
)

const (
	ApplicationJsonHeader = "application/json"
	CertsFilePermissions = 0644
)

const (
	CertLetsEncryptStatement = "/etc/letsencrypt/live/%s/fullchain.pem"
	KeyLetsEncryptStatement = "/etc/letsencrypt/live/%s/privkey.pem"
	SandboxCertUrlStatement = "https://%s/sandbox"
)

// certbot certonly --standalone --non-interactive --agree-tos --email you.are@awesome.com --domains awesome.com,www.awesome.com
const (
	Certbot = "certbot"
	CertOnly = "certonly"
	StandaloneOption = "--standalone"
	NonInteractiveOption = "--non-interactive"
	AgreeTosOption = "--agree-tos"
	EmailOption = "--email"
	DomainsOption = "--domains"
)

// $ openssl req -x509 -sha256 -newkey rsa:2048 -keyout certificate.key -out certificate.crt -days 1024 -nodes -subj "/C=<Country Code>/ST=<State>/L=<City>/O=<Organization>/CN=<Common Name>"
const (
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
	
	// Addresses
	AuthnAddress = os.Getenv("AUTHN_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_ADDRESS")
	MailAddress = os.Getenv("MAIL_ADDRESS")
	MediaAddress = os.Getenv("MEDIA_ADDRESS")
	LogsAddress = os.Getenv("LOGS_ADDRESS")

	// domains for certificates
	ProductionHostname = os.Getenv("PRODUCTION_HOSTNAME")
	ProductionDomains = os.Getenv("PRODUCTION_DOMAINS")

	SandboxHostname = os.Getenv("SANDBOX_HOSTNAME")
	SandboxDomains = os.Getenv("SANDBOX_DOMAINS")
	SandboxUser = os.Getenv("SANDBOX_USER")
	SandboxPassword = os.Getenv("SANDBOX_PASSWORD")

	// certificates
	CertbotEmail = os.Getenv("CERTBOT_EMAIL")
	CertFilepath = os.Getenv("CERTS_CRT_FILEPATH")
	KeyFilepath = os.Getenv("CERTS_KEY_FILEPATH")
	CountryCode = os.Getenv("CERTS_COUNTRY_CODE")
	State = os.Getenv("CERTS_STATE")
	City = os.Getenv("CERTS_CITY")
	Organization = os.Getenv("CERTS_ORGANIZATION")

	ProductionCertLetsEncryptFilepath = GetStatement(CertLetsEncryptStatement, ProductionHostname)
	ProductionKeyLetsEncryptFilepath = GetStatement(KeyLetsEncryptStatement, ProductionHostname)
	
	SandboxCertLetsEncryptFilepath = GetStatement(CertLetsEncryptStatement, SandboxHostname)
	SandboxKeyLetsEncryptFilepath = GetStatement(KeyLetsEncryptStatement, SandboxHostname)
	
	SandboxCertUrlAddress = GetStatement(SandboxCertUrlStatement, ProductionHostname)
	
	SelfSignedSubject = GetSelfSignedCertificateSubjectStatement()
)

func GetStatement(statement string, compliment string) (string) {
	return fmt.Sprintf(statement, compliment)
}

func GetSelfSignedCertificateSubjectStatement() string {
	return fmt.Sprintf(
		SubjectStatement,
		CountryCode,
		State,
		City,
		Organization,
		ProductionHostname,
	)
}