//	brian taylor vann
//	briantaylorvann dot com

package server

import (
	"net/http"
	
	"webapi/routes"
)

const httpPort  = ":80"

func Create() {
	mux := routes.CreateMux()

	http.ListenAndServe(
		httpPort,
		mux,
	)
}
