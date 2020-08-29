//	brian taylor vann
//	briantaylorvann dot com

package main

import (
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"webapi/server"
)

func main() {
	sessionx.Setup()

	server.Create()
}
