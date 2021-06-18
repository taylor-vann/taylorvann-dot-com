package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/mux"
)

var (
	httpPort = fmt.Sprint(":", details.Details.Server.HTTPPort)
)

func main() {
	httpMux := mux.CreateMux()

	http.ListenAndServe(
		httpPort,
		httpMux,
	)
}
