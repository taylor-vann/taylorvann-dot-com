//	brian taylor vann
//	briantaylorvann dot com

package server

import (
	"net/http"
	"os"
	
	"webapi/server/muxrouter"
)

const (
	httpPort  = ":80"
	httpsPort = ":443"
)

var (
	certFilepath = os.Getenv("CERTS_CRT_FILEPATH")
	keyFilepath = os.Getenv("CERTS_KEY_FILEPATH")
)

func CreateServer() {
	proxyMux := muxrouter.CreateProxyMux()
	mux := muxrouter.CreateRedirectToHttpsMux()

	go http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)

	http.ListenAndServe(
		httpPort,
		mux,
	)
}
