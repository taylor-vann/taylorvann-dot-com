//	brian taylor vann

package mux

import (
	"encoding/json"
	"net/http"

	"webapi/details"
	"webapi/setterx"
)

const (
	contentType = "Content-Type"
	applicationJson = "application/json"
	detailsRoute = "/details"
	detailsRouteWithSlash = "/details"
	getRoute   = "/get"
	getRouteWithSlash   = "/get/"
	setRoute   = "/set"
	setRouteWithSlash   = "/set/"
)

type GetEntryRequestBody = string

type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

type ErrorDeclarations = []ErrorEntity

const (
	incorrectRequest = "request body structure is incorrect"
	failedToGet = "failed to return entry"
	failedToSet = "failed to set address and entry"
)

var (
	setter, errSetter = setterx.Create(&details.Details.Cache)
)

func writeError(w http.ResponseWriter, kind string, message string) {
	setErrors := ErrorDeclarations{
		ErrorEntity{
			Kind: kind,
			Message: message,
		},
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(setErrors)
	
}

func writeGetEntry(w http.ResponseWriter, entry interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(entry)
}

func writeSetEntry(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)
}

func setEntry(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		writeError(w, incorrectRequest, "nil request body found")
		return
	}

	var rBody setterx.SetBody
	errRBody := json.NewDecoder(r.Body).Decode(&rBody)
	if errRBody != nil {
		writeError(w, incorrectRequest, errRBody.Error())
		return
	}

	_, errEntry := setter.Set(&rBody)
	if errRBody != nil {
		writeError(w, failedToSet, errEntry.Error())
		return
	}

	writeSetEntry(w)
}

func getEntry(w http.ResponseWriter, r *http.Request) {
	var address string
	errRBody := json.NewDecoder(r.Body).Decode(&address)
	if errRBody != nil {
		writeError(w, incorrectRequest, errRBody.Error())
		return
	}

	if r.Body == nil {
		writeError(w, incorrectRequest, errRBody.Error())
		return
	}

	entry, errEntry := setter.Get(address)
	if errRBody != nil {
		writeError(w, failedToGet, errEntry.Error())
		return
	}

	writeGetEntry(w, entry);
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(setRoute, getEntry)
	mux.HandleFunc(setRouteWithSlash, getEntry)

	mux.HandleFunc(getRoute, setEntry)
	mux.HandleFunc(getRouteWithSlash, setEntry)

	return mux
}
