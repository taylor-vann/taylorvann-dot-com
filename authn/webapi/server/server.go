package server

import (
	"net/http"
	"webapi/routes"
)

func CreateServer(port int) {
	muxHttps := http.NewServeMux()
	routes.CreateRoutes(muxHttps)

	http.ListenAndServeTLS(
		":5000",
		"/usr/local/certs/authn/https-server.crt",
		"/usr/local/certs/authn/https-server.key",
		muxHttps,
	)
}
