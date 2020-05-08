package main

import (
	"webapi/certs"
	"webapi/mailx"
	"webapi/server"
)

func main() {
	certs.Create()
	mailx.Setup()
	server.CreateServer()
}
