//	brian taylor vann
//	gateway

package main

import (
	"net/http"
	
	"webapi/details"
	"webapi/muxrouter"
)

const (
	httpPort  = ":80"
	httpsPort = ":443"
)

var (
	certFilepath = details.Details.CertPaths.Cert
	keyFilepath  = details.Details.CertPaths.PrivateKey
)

func main() {
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
