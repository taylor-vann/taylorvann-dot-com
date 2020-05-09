package server

import (
	"net/http"

	"webapi/routes"
	"webapi/server/constants"
	certsConstants "github.com/taylor-vann/tvgtb/certificatesx/constants"
)

func CreateServer() {
	muxHttps := http.NewServeMux()
	routes.CreateRoutes(muxHttps)

	http.ListenAndServeTLS(
		constants.Ports.Https,
		certsConstants.Filepaths.Cert,
		certsConstants.Filepaths.Key,
		muxHttps,
	)
}
