//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Serve single page app
	mux.HandleFunc("/", fileserver.ServeHomeApp)
	mux.HandleFunc("/scripts/", fileserver.ServeHomeFiles)
	mux.HandleFunc("/styles/", fileserver.ServeHomeFiles)

	mux.HandleFunc("/sign-in/", fileserver.ServeSignInFiles)

	mux.HandleFunc("/internal/", fileserver.ServeInternalApp)
	mux.HandleFunc("/internal/scripts/", fileserver.ServeInternalFiles)
	mux.HandleFunc("/internal/styles/", fileserver.ServeInternalFiles)

	return mux
}
