//	brian taylor vann
//	taylorvann dot com

//  Routes Gateway
//	Keep routes separate and isolated for easier scaling.
//	Each route could potentially be replaced by a simple
//	http request to an external service.

// Package routes -
package routes

import (
	"net/http"
	"webapi/hooks/documents"
	"webapi/hooks/scripts"
	"webapi/hooks/styles"
)

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	//	documents
	mux.HandleFunc("/", documents.Homepage)

	// supplementary, need any jwt
	mux.HandleFunc("/styles/", styles.Styles)
	mux.HandleFunc("/scripts/", scripts.Scripts)

	//	request session, needs guest jwt
	mux.HandleFunc("/login/", documents.Login)
	mux.HandleFunc("/logout/", documents.Logout)

	return mux
}
