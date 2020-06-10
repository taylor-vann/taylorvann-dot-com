package main

import (
	"webapi/server"
	"webapi/store"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx"
)

func main() {	
	store.CreateRequiredTables()
	store.InitFromJSON()

	server.CreateServer()

	clientx.Setup()
}
