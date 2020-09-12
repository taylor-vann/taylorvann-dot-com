//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"
	"os"

	"webapi/fileserver"
)

const Environment = os.Getenv("STAGE")

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

	// development serve tests

	if Environment == "DEVELOPMENT" {
		mux.HandleFunc("/tests/", fileserver.ServeInternalFiles)
		mux.HandleFunc("/internal/tests/", fileserver.ServeInternalFiles)	
	}
	
	return mux
}
