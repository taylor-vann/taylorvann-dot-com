//	brian taylor vann
//	taylorvann dot com

// 	Authn - Routes / Gateway


package routes

import (
	"net/http"
	"webapi/routes/ping"
	"webapi/sendmail"
)

func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", ping.Details)
	mux.HandleFunc("/sendmail", sendmail.NoReply)

	return mux
}
