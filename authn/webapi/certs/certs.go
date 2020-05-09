package certs

import (
	"os/exec"
)

func runScript(scriptName string) {
	cmd := exec.Command("/bin/sh", scriptName)
	_, err := cmd.Output()

	if err != nil {
		return
	}
}

func Create() {
	runScript("/usr/local/certs/generate_self_signed_certificate.sh")
}
