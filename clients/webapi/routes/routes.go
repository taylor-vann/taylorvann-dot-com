//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Serve internal files
	mux.HandleFunc("/sign-in/", fileserver.ServeSignIn)

	// Serve files
	mux.HandleFunc("/", fileserver.ServeLanding)

	return mux
}
