//	brian taylor vann
//	briantaylorvann dot com

package server

import (
	"net/http"
	
	"webapi/certificatesx/constants"
	"webapi/server/routes"
)

const (
	Http  = ":80"
	Https = ":443"
)

func CreateServer() {
	proxyMux := routes.CreateProxyMux()
	mux := routes.RedirectToHttpsMux()

	go http.ListenAndServeTLS(
		Https,
		constants.CertFilepath,
		constants.KeyFilepath,
		proxyMux,
	)

	http.ListenAndServe(
		Http,
		mux,
	)
}
