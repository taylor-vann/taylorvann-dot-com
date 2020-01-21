package server

import (
	"net/http"
	"webapi/routes"
)

// CreateServer - Start a websever
func CreateServer(port int) {
	mux := http.NewServeMux()

	// send mux to local routes package to append hooks
	routes.CreateRoutes(mux)

	// start app
	http.ListenAndServe(":5000", mux)
}
