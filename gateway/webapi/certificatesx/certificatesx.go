// brian taylor vann
// briantaylorvann dot com

package certificatesx

import (
	"os/exec"

	"webapi/certificatesx/constants"
)

// $ openssl req -new -x509 -sha256 -newkey rsa:2048 -keyout certificate.key -out certificate.crt -days 356 -nodes -subj "/C=<Country Code>/ST=<State>/L=<City>/O=<Organization>/CN=<Common Name>"
func execSelfSignedCertificateScript() (string, error) {
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

func Create() (string, error) {
	if constants.Environment == "PRODUCTION" {
		// production certificate should already be present
		// read if files are present
		// error if they are not
		return "", nil
	}
	
	if constants.Environment == "DEVELOPMENT" {
		// fetch sandbox certificates from production
		// return "", FetchSandboxCertificates()
		return "", nil
	}

	return execSelfSignedCertificateScript()
}