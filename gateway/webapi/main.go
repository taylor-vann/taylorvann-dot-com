//	brian taylor vann
//	briantaylorvann dot com

package main

import (
	"webapi/certificatesx"
	"webapi/server"
)

func main() {
	output, _ := certificatesx.Create()

	server.CreateServer()
}
