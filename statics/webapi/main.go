package main

import (
	"webapi/certs"
	"webapi/server"
)

func main() {
	certs.Create()
	server.CreateServer()
}
