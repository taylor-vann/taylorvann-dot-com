package routes

import (
	"net/http"

	"webapi/hooks"
)

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	// ping - information endpoint about our authentication api
	mux.HandleFunc("/", hooks.Ping)

	// personal accounts - endpoints regarding personal details
	mux.HandleFunc("/create_user", hooks.CreateUser)
	mux.HandleFunc("/update_account_details", hooks.Ping)
	mux.HandleFunc("/update_password", hooks.Ping)
	mux.HandleFunc("/remove_user", hooks.Ping)

	// validation - arbitrarily validate a user's password
	mux.HandleFunc("/validate_user", hooks.Ping)

	// sessions - mux.HandleFunc("/create_session", hooks.Ping)
	mux.HandleFunc("/invalidate_session", hooks.Ping)
	// our most frequently called method
	mux.HandleFunc("/return_session", hooks.Ping)

	return mux
}
