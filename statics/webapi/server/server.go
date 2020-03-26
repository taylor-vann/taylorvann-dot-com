package server

import (
	"net/http"
	"webapi/routes"
)

// CreateServer - Start a websever
func CreateServer() {
	mux := http.NewServeMux()

	// send mux to local routes package to append hooks
	routes.CreateRoutes(mux)

	// start app
	http.ListenAndServeTLS(
		":4000",
		"/usr/local/certs/local_gateway/https-server.crt",
		"/usr/local/certs/local_gateway/https-server.key",
		mux,
	)
}