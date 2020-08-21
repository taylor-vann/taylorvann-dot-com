//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// log out
	// request quest cookie
	// login
	mux.HandleFunc("/", fileserver.Serve)

	return mux
}
