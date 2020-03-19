package server

import (
	"net/http"
	"webapi/routes"
	"webapi/store"
)

// CreateServer - Start a websever
func CreateServer(port int) {
	// create databases
	store.CreateRequiredTables()
	store.InitFromJSON()

	// send mux to local routes package to append hooks
	mux := http.NewServeMux()
	routes.CreateRoutes(mux)

	// start app
	http.ListenAndServe(":5000", mux)
}
