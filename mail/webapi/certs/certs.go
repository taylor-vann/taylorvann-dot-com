// brian taylor vann
// taylorvann dot com

package certs

import (
	"os/exec"

	"webapi/certs/constants"
)

func runScript(script string) {
	cmd := exec.Command("/bin/sh", script)
	_, err := cmd.Output()

	if err != nil {
		return
	}
}

func Create() {
	runScript(constants.Addresses.Script)
}
