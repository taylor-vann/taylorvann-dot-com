//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Serve static files
	mux.HandleFunc("/scripts/", fileserver.Serve)
	mux.HandleFunc("/styles/", fileserver.Serve)

	// Otherwise serve application
	mux.HandleFunc("/", fileserver.ServeLanding)

	return mux
}
