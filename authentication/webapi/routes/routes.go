//	brian taylor vann
//	taylorvann dot com

//	Keep routes separate and isolated for easier scaling.
//	Each route could potentially be replaced by a simple
//	http request to an external service.

// Package routes -
package routes

import (
	"net/http"
	"webapi/hooks/ping"
	"webapi/hooks/sessions"
	"webapi/hooks/store"
)

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	//	ping
	mux.HandleFunc("/", ping.Details)

	//	store
	mux.HandleFunc("/store/q/", store.Query)
	mux.HandleFunc("/store/m/", store.Mutation)

	//	sessions
	mux.HandleFunc("/sessions/q/", sessions.Query)
	mux.HandleFunc("/sessions/m/", sessions.Mutation)

	return mux
}
