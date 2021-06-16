//	brian taylor vann

package muxrouter

import (
	"encoding/json"
	"net/http"

	"webapi/setterx"
)

const (
	getDetails = "/details"
	getRoute   = "/get"
	getRouteWithSlash   = "/get/"
	setRoute   = "/set"
	setRouteWithSlash   = "/set/"
)

type SetRequestBody struct {
	Address string      `json:"address"`
	Entry   interface{} `json:"entry"`
}

type GetEntryRequestBody = string

type ErrorEntity struct {
	kind    string `json:"kind"`
	message string `json:"message"`
}

type ErrorResponse = []ErrorEntity

const (
	incorrectRequest = "request body structure is incorrect"
	failedToGet = "failed to return entry"
	failedToSet = "failed to set address and entry"

)

func writeError(w http.ResponseWriter, kind string, message string) {
	setErrors := ErrorResponse{
		ErrorEntity{
			kind,
			message,
		},
	}
	json.NewEncoder(w).Encode(setErrors)
	w.WriteHeader(http.StatusBadRequest)
}

func writeGetEntry(w http.ResponseWriter, entry interface{}) {
	json.NewEncoder(w).Encode(entry)
	w.WriteHeader(http.StatusOK)
}

func writeSetEntry(w http.ResponseWriter, entry interface{}) {
	w.WriteHeader(http.StatusOK)
}

func setEntry(w http.ResponseWriter, r *http.Request) {
	var rBody SetRequestBody
	errRBody := json.NewDecoder(r.Body).Decode(&rBody)
	if errRBody != nil {
		writeError(w, incorrectRequest, errRBody.Error())
		return
	}

	entry, errEntry := setterx.Set(address)
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

	entry, errEntry := setterx.Get(address)
	if errRBody != nil {
		writeError(w, failedToGet, errEntry.Error())
		return
	}

	writeGetEntry(w, entry);
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(setRoute, getEntry)
	mux.HandleFunc(setRoute, getEntryWithSlash)

	mux.HandleFunc(getRoute, setEntry)
	mux.HandleFunc(getRoute, setEntryWithSlash)

	return mux
}
