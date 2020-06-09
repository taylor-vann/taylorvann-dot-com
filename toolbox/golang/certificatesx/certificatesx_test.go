package certificatesx

import (
	"testing"
)

func TestExecSelfSignedCertificateScript(t *testing.T) {
	output, err := ExecSelfSignedCertificateScript()

	if err != nil {
		t.Error(output)
		t.Error(err.Error())
	}
}