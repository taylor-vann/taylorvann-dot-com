//	brian taylor vann
//	taylorvann dot com

package server

import (
	"net/http"
	
	"webapi/routes"

	constants "webapi/server/constants"
)

func CreateServer() {
	proxyMux := routes.CreateProxyMux()
	mux := routes.RedirectToHttpsMux()

	go http.ListenAndServeTLS(
		constants.Https,
		constants.CertFilepath,
		constants.KeyFilepath,
		proxyMux,
	)

	http.ListenAndServe(
		constants.Http,
		mux,
	)
}
