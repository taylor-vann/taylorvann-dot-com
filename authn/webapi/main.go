package main

import (
	"webapi/server"
	"webapi/store"
)

func main() {
	store.CreateRequiredTables()
	store.InitFromJSON()

	server.CreateServer()
}
