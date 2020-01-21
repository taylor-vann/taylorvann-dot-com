package main

import (
	"webapi/server"
)

func main() {
	// read json environment here

	// create server
	server.CreateServer(5000)
}
