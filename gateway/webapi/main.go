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
	proxyMux, errProxyMux := muxrouter.CreateProxyMux(&details.Details.Routes)
	if (errProxyMux != nil) {
		return
	}

	go http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)

	mux := muxrouter.CreateRedirectToHttpsMux()
	http.ListenAndServe(
		httpPort,
		mux,
	)
}
