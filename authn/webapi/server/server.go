package server

import (
	"net/http"
	"webapi/routes"

	certsConstants "github.com/taylor-vann/tvgtb/certificatesx/constants"
)

func CreateServer(port int) {
	muxHttps := http.NewServeMux()
	routes.CreateRoutes(muxHttps)

	http.ListenAndServeTLS(
		":5000",
		certsConstants.Filepaths.Cert,
		certsConstants.Filepaths.Key,
		muxHttps,
	)
}
