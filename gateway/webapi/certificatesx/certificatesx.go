// brian taylor vann
// taylorvann dot com

package certificatesx

import (
	"os/exec"

	"webapi/certificatesx/constants"
)

// write file to disk

// $ openssl req -new -x509 -sha256 -newkey rsa:2048 -keyout certificate.key -out certificate.crt -days 356 -nodes -subj "/C=<Country Code>/ST=<State>/L=<City>/O=<Organization>/CN=<Common Name>"
func ExecSelfSignedCertificateScript() (string, error) {
	cmd := exec.Command(
		constants.OpenSSL,
		constants.Req,
		constants.NewOption,
		constants.X509Option,
		constants.Sha256Option,
		constants.NewKeyOption,
		constants.RSAEncoding,
		constants.DaysOption,
		constants.ThreeSixtyFive,
		constants.NodesOption,
		constants.CertOutOption,
		constants.CertFilepath,
		constants.KeyOutOption,
		constants.KeyFilepath,
		constants.SubjectOption,
		constants.SelfSignedSubject,
	)

	output, err := cmd.CombinedOutput()
	return string(output), err
}

// certbot certonly --standalone --non-interactive --agree-tos --email <email> --domains <domain,domain2,...>
func ExecCertbotProductionCertificates() (string, error) {
	cmd := exec.Command(
		constants.Certbot,
		constants.StandaloneOption,
		constants.NonInteractiveOption,
		constants.AgreeTosOption,
		constants.EmailOption,
		constants.CertbotEmail,
		constants.DomainsOption,
		constants.ProductionDomains,
	)

	output, err := cmd.Output()
	return string(output), err
}

// certbot certonly --standalone --non-interactive --agree-tos --email <email> --domains <domain,domain2,...>
func ExecCertbotSandboxCertificates(domains string) (string, error) {
	cmd := exec.Command(
		constants.Certbot,
		constants.StandaloneOption,
		constants.NonInteractiveOption,
		constants.AgreeTosOption,
		constants.EmailOption,
		constants.CertbotEmail,
		domains,
		constants.SandboxDomains,
	)

	output, err := cmd.Output()
	return string(output), err
}

func Create() (string, error) {
	if constants.Environment == "PRODUCTION" {
		outputProdCert, errProdCert := ExecCertbotProductionCertificates()
		if errProdCert != nil {
			return outputProdCert, errProdCert
		}

		return ExecCertbotSandboxCertificates(constants.DomainsOption)
	}

	if constants.Environment == "DEVELOPMENT" {
		// fetch sandbox certificates from production
		// return "", FetchSandboxCertificates()
	}

	return ExecSelfSignedCertificateScript()
}