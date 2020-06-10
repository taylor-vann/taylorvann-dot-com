//	brian taylor vann
//	briantaylorvann dot com

package server

import (
	"net/http"
	
	"webapi/routes"
)

const Http  = ":80"

func Create() {
	mux := routes.Create()

	http.ListenAndServe(
		Http,
		mux,
	)
}
