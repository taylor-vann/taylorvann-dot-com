//	brian taylor vann
//	taylorvann dot com

// 	Authn - Routes / Gateway


package routes

import (
	"net/http"
	"webapi/routes/ping"
	"webapi/sendonly"
)

func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", ping.Details)
	mux.HandleFunc("/sendonly/", sendonly.NoReply)

	return mux
}
