//	brian taylor vann
//	gateway

package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/mux"
)

var (
	httpsPort = fmt.Sprint(":", details.Details.Server.HTTPSPort)
	certFilepath = details.Details.CertPaths.Cert
	keyFilepath  = details.Details.CertPaths.PrivateKey
)

func main() {
	httpMux := mux.CreateMux()

	http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		httpMux,
	)
}
