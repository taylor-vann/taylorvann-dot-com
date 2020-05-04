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
	usersHooks "webapi/store/users/hooks"
	rolesHooks "webapi/store/roles/hooks"
)

func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", ping.Details)

	mux.HandleFunc("/q/users/", usersHooks.Query)
	mux.HandleFunc("/m/users/", usersHooks.Mutation)

	mux.HandleFunc("/q/roles/", rolesHooks.Query)
	mux.HandleFunc("/m/roles/", rolesHooks.Mutation)
		
	mux.HandleFunc("/q/sessions/", sessionHooks.Query)
	mux.HandleFunc("/m/sessions/", sessionHooks.Mutation)

	return mux
}
