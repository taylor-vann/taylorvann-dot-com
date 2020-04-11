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
	// "webapi/hooks/store"
)

//	CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	//	ping
	mux.HandleFunc("/", ping.Details)

	// //	users
	// mux.HandleFunc("/q/users/", store.Query)
	// mux.HandleFunc("/m/users/", store.Mutation)

	// //	roles
	// mux.HandleFunc("/q/roles/", store.Query)
	// mux.HandleFunc("/m/roles/", store.Mutation)
		
	//	sessions
	mux.HandleFunc("/q/sessions/", sessions.Query)
	mux.HandleFunc("/m/sessions/", sessions.Mutation)

	return mux
}
