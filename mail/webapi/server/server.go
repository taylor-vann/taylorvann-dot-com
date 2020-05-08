package server

import (
	"net/http"

	"webapi/routes"
	"webapi/server/constants"
	certsConstants "webapi/certs/constants"
)

func CreateServer() {
	muxHttps := http.NewServeMux()
	routes.CreateRoutes(muxHttps)

	http.ListenAndServeTLS(
		constants.Ports.Https,
		certsConstants.Addresses.Cert,
		certsConstants.Addresses.Key,
		muxHttps,
	)
}
