package main

import (
	"webapi/server"
	"webapi/store"

	"github.com/taylor-vann/tvgtb/certificatesx"
)

func main() {
	certificatesx.Create()
	store.CreateRequiredTables()
	server.CreateServer(5000)
}
