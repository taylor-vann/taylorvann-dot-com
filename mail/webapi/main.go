package main

import (
	"webapi/mailbox/mailx"
	"webapi/server"

	"github.com/taylor-vann/tvgtb/certificatesx"
)

func main() {
	certificatesx.Create()
	mailx.Setup()
	server.CreateServer()
}
