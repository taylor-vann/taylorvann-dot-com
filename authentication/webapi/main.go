package main

import (
	"webapi/server"
	"webapi/whitelist"
	"webapi/store"
)

func main() {
	// create database if necessary
	store.CreateRequiredDatabases()

	// ping our redis store
	whitelist.HelloWorld()
	
	// create server
	server.CreateServer(5000)
}
