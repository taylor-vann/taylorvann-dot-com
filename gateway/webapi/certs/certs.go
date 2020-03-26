package certs

import (
	"os/exec"

	"webapi/constants"
)

func Create() {
	cmd := exec.Command("/bin/sh", constants.CertAddresses.Script)
	_, err := cmd.Output()

	if err != nil {
		// chance to log output to log service
		return
	}
}
