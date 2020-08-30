//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/server/details"
	sessionHooks "webapi/sessions/hooks"
	rolesHooks "webapi/store/roles/hooks"
	usersHooks "webapi/store/users/hooks"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/details/", details.Details)

	mux.HandleFunc("/q/users/", usersHooks.Query)
	mux.HandleFunc("/m/users/", usersHooks.Mutation)

	mux.HandleFunc("/q/roles/", rolesHooks.Query)
	mux.HandleFunc("/m/roles/", rolesHooks.Mutation)

	mux.HandleFunc("/q/sessions/", sessionHooks.Query)
	mux.HandleFunc("/m/sessions/", sessionHooks.Mutation)

	return mux
}
