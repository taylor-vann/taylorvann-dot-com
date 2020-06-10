package server

import (
	"net/http"

	"webapi/routes"
	"webapi/server/constants"
	certsConstants "github.com/taylor-vann/weblog/authn/toolbox/certificatesx/constants"
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
