package server

import (
	"net/http"

	"webapi/server/routes"
)

const (
	Http = ":80"
)

func CreateServer() {
	mux := routes.CreateMux()

	http.ListenAndServe(
		Http,
		mux,
	)
}
