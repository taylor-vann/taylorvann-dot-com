package server

import (
	"net/http"
	
	"webapi/constants"
	"webapi/routes"
)

// CreateServer - Start a websever
func CreateServer() {
	proxyMux := routes.CreateProxyMux()
	mux := routes.RedirectToHttpsMux()

	go http.ListenAndServeTLS(
		constants.Ports.Https,
		constants.CertAddresses.Cert,
		constants.CertAddresses.Key,
		proxyMux,
	)

	http.ListenAndServe(
		constants.Ports.Http,
		mux,
	)
}
