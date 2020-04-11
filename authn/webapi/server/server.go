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
	// store.InitFromJSON()

	// send mux to local routes package to append hooks
	muxHttps := http.NewServeMux()
	routes.CreateRoutes(muxHttps)

	// start app
	http.ListenAndServeTLS(
		":5000",
		"/usr/local/certs/authn/https-server.crt",
		"/usr/local/certs/authn/https-server.key",
		muxHttps,
	)
}
