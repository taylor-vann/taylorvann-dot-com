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
	"webapi/routes/ping"
	sessionHooks "webapi/sessions/hooks"
	// "webapi/sessions/hooks/users"

	// "webapi/hooks/users"
	// "webapi/hooks/roles"
)

//	CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	//	ping
	mux.HandleFunc("/", ping.Details)

	//	users
	// mux.HandleFunc("/q/users/", users.Query)
	// mux.HandleFunc("/m/users/", users.Mutation)

	// //	roles
	// mux.HandleFunc("/q/roles/", roles.Query)
	// mux.HandleFunc("/m/roles/", roles.Mutation)
		
	//	sessions
	mux.HandleFunc("/q/sessions/", sessionHooks.Query)
	mux.HandleFunc("/m/sessions/", sessionHooks.Mutation)

	return mux
}
