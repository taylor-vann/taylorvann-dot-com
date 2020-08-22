//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

		// Issue and Destroy sessions without file-serving
	// mux.HandleFunc("/request_session")
	// mux.HandleFunc("/remove_session")

	// Serve internal files
	mux.HandleFunc("/sign-in/", fileserver.ServeSignIn)
	mux.HandleFunc("/", fileserver.Serve)

	return mux
}
