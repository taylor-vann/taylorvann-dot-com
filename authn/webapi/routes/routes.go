//	brian taylor vann
//	taylorvann dot com

// 	Authn - Routes / Gateway
//
//	Keep routes separate and isolated for easier scaling.
//	Each route should be potentially replaced by a simple
//	http request to an external service.

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
