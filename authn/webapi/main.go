package main

import (
	"webapi/certs"
	"webapi/server"
	"webapi/store"
)

func main() {
	certs.Create()
	store.CreateRequiredTables()
	server.CreateServer(5000)
}
