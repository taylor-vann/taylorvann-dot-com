//	brian taylor vann

package muxrouter

import (
	"net/http"
)

const (
	getDetails	= "/details"
	getRoute		= "/get"
	setRoute		= "/set"
)

// json request
type LimiterRequestPayload struct {
	Address	string	`json:"address"`
}

func setEntry(w http.ResponseWriter, r *http.Request) {}

func getEntry(w http.ResponseWriter, r *http.Request) {}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(getRoute, setEntry)
	mux.HandleFunc(setRoute, getEntry)

	return mux
}
