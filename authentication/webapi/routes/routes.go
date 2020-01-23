package routes

import (
	"encoding/json"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world, json!")
}

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", helloWorld)
	mux.HandleFunc("/session", helloWorld)
	mux.HandleFunc("/user", helloWorld)

	return mux
}