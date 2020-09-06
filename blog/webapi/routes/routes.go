//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/details"
	"webapi/sessions"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/details/", details.Details)

	// Sessions
	mux.HandleFunc("/sessions/create_guest_session/", sessions.RequestGuestSession)
	mux.HandleFunc("/sessions/create_client_session/", sessions.RequestClientSession)
	mux.HandleFunc("/sessions/delete_session/", sessions.RequestGuestSession)

	return mux
}
