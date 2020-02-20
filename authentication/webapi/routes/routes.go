package routes

import (
	"net/http"
	// "webapi/hooks"
)

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	// ping - information endpoint about our authentication api
	// mux.HandleFunc("/", hooks.Ping)

	// /store/m
	// /store/q
	// check cache for stale request, return stale receipt
	// set query to pending
	// query store
	// save in cache
	// mux.HandleFunc("/store/", hooks.Ping)

	// /q
	// read a cache
	// query store if nothing in cache
	// store results in cache
	// we don't count sessions as mutations because
	// our sessions are supposed to be stateless
	// mux.HandleFunc("/sessions/", hooks.Ping)

	return mux
}
