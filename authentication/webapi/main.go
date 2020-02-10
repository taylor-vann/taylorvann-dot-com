package main

import (
	"webapi/constants"
	"webapi/server"
	// "webapi/store"
	"webapi/whitelist"
)

func main() {
	// create database if necessary
	// store.CreateRequiredDatabases()
	constants.SetEnvironmentConstants()

	// ping our redis store
	whitelist.HelloWorld()

	// create server
	server.CreateServer(5000)
}
