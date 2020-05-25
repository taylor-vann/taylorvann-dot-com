//	brian taylor vann
//	briantaylorvann dot com

// 	Authn - Routes / Gateway


package routes

import (
	"net/http"
	"webapi/routes/ping"
	"webapi/mailbox/hooks"
)

func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", ping.Details)
	mux.HandleFunc("/sendonly/", hooks.NoReply)

	return mux
}
