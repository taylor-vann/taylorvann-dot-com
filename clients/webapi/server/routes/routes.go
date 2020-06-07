//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"
	
	"webapi/server/routes/ping"
	"webapi/server/routes/home"
	// "webapi/server/routes/internal"
	// "webapi/server/routes/login"
	// "webapi/server/routes/logout"
)

func Create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping.Details)
	mux.HandleFunc("/", home.ServeFiles)
	// mux.HandleFunc("/login", login.Template)
	// mux.HandleFunc("/logout", logout.Template)
	// mux.HandleFunc("/internal", usersHooks.Template)

	return mux
}
