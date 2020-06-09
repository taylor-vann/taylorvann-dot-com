// brian taylor vann
// toolbox-go

package certificatesx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os/exec"
	"io/ioutil"

	"github.com/weblog/toolbox/golang/certificatesx/constants"
	"github.com/weblog/toolbox/golang/certificatesx/requests"
	"github.com/weblog/toolbox/golang/certificatesx/responses"
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
func ExecCertbotSandboxCertificates() (string, error) {
	cmd := exec.Command(
		constants.Certbot,
		constants.StandaloneOption,
		constants.NonInteractiveOption,
		constants.AgreeTosOption,
		constants.EmailOption,
		constants.CertbotEmail,
		constants.DomainsOption,
		constants.SandboxDomains,
	)

	output, err := cmd.Output()
	return string(output), err
}

func WriteSandboxCertificatesToDisk(certs *responses.Certificates) error {
	certBytes, errCertBytes := json.Marshal(certs.Cert)
	if errCertBytes != nil {
		return errCertBytes
	}
	errWriteCert := ioutil.WriteFile(
		constants.SandboxCertLetsEncryptFilepath,
		certBytes,
		constants.CertsFilePermissions,
	)
	if errWriteCert != nil {
		return errWriteCert
	}

	keyBytes, errKeyBytes := json.Marshal(certs.Key)
	if errKeyBytes != nil {
		return errKeyBytes
	}
	errWriteKey := ioutil.WriteFile(
		constants.SandboxKeyLetsEncryptFilepath,
		keyBytes,
		constants.CertsFilePermissions,
	)
	return errWriteKey
}

func FetchSandboxCertificates() (error) {	
	body := requests.Body{
		Action: "REQUEST_SANDBOX_CERTIFICATES",
		Params: &requests.Params{
			Email: constants.SandboxUser,
			Password: constants.SandboxPassword,
		},
	}

	bodyJsonBuffer := new(bytes.Buffer)
	json.NewEncoder(bodyJsonBuffer).Encode(body)
	resp, errResp := http.Post(
		constants.SandboxCertUrlAddress,
		constants.ApplicationJsonHeader,
		bodyJsonBuffer,
	)
	if errResp != nil {
		return errResp
	}
	defer resp.Body.Close()
	if resp.Body == nil {
		errors.New("nil body returned")
	}

	var responseBody responses.Body
	errJsonResponse := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJsonResponse != nil {
		return errJsonResponse
	}
	if responseBody.Certificates != nil {
		return WriteSandboxCertificatesToDisk(responseBody.Certificates)
	}

	return errors.New("failed to retrieve sandbox certificates")
}

func Create() (string, error) {
	if constants.Environment == "PRODUCTION" {
		outputProdCert, errProdCert := ExecCertbotProductionCertificates()
		if errProdCert != nil {
			return outputProdCert, errProdCert
		}

		return ExecCertbotSandboxCertificates()
	}

	if constants.Environment == "DEVELOPMENT" {
		// fetch sandbox certificates from production
		return "", FetchSandboxCertificates()
	}

	return ExecSelfSignedCertificateScript()
}