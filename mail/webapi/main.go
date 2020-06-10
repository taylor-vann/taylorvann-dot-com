package main

import (
	"webapi/mailbox/mailx"
	"webapi/server"

	"github.com/taylor-vann/weblog/authn/toolbox/certificatesx"
)

func main() {
	certificatesx.Create()
	mailx.Setup()
	server.CreateServer()
}
