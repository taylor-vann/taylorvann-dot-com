package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/muxrouter"
)

var (
	httpPort = fmt.Sprint(":", details.Details.Server.HTTPPort)
)

func main() {
	httpMux := muxrouter.CreateMux()

	http.ListenAndServe(
		httpPort,
		httpMux,
	)
}
