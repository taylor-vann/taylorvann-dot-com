//	brian taylor vann
//	gateway

package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/muxrouter"
)

var (
	httpPort     = fmt.Sprint(":", details.Details.Server.HTTPPort)
	httpsPort    = fmt.Sprint(":", details.Details.Server.HTTPSPort)
	certFilepath = details.Details.CertPaths.Cert
	keyFilepath  = details.Details.CertPaths.PrivateKey
)

func main() {
	proxyMux, errProxyMux := muxrouter.CreateHTTPSMux(&details.Details.Routes)
	if errProxyMux != nil {
		return
	}

	go http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)

	mux := muxrouter.CreateRedirectMux()
	http.ListenAndServe(
		httpPort,
		mux,
	)
}
