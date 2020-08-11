//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"
	"os"
	"strings"

	"webapi/routes/ping"
)

var (
	webClientsDirectory = os.Getenv("WEB_CLIENTS_DIRECTORY")
	
	directoryRune = []byte("/")[0]
	relativeRune = []byte(".")[0]

	fileServer = http.FileServer(http.Dir(webClientsDirectory))
)

func isValidPath(path string) bool {
	searchIndex := 0
	pathLength := len(path)
	for searchIndex < pathLength {
		searchIndex += 1
		if path[searchIndex] == relativeRune && path[searchIndex - 1] == directoryRune {
			return false
		}
	}

	return true
}

// Blog Location
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check for correct paths
	if isValidPath(r.URL.Path) {
		fileServer.ServeHTTP(w, r)
	}

	// default 404
	w.WriteHeader(http.StatusNotFound)
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", ping.Details)
	mux.HandleFunc("/", serveStaticFiles)

	return mux
}
