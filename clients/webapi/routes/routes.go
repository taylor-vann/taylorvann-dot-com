//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"

	"webapi/fileserver"
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	// briantaylorvann.com/home/ or briantaylorvann.com/
	// for single page app snugness

	// Serve static files
	mux.HandleFunc("/", fileserver.Serve)

	return mux
}
