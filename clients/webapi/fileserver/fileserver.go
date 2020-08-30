package fileserver

import (
	"net/http"
	"os"
)

var (
	webClientsDirectory    = os.Getenv("WEB_CLIENTS_DIRECTORY")
	waywardRequestFilename = webClientsDirectory + "/lost/"

	relativeRune = []byte(".")[0]
)

func containsRelativeBackPaths(path string) bool {
	pathLength := len(path)

	searchIndex := 1
	for searchIndex < pathLength {
		if path[searchIndex] == relativeRune &&
			path[searchIndex-1] == relativeRune {
			return true
		}
		searchIndex += 1
	}

	return false
}

func serveWaywardRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, waywardRequestFilename)
}

func serveStaticFiles(
	w http.ResponseWriter,
	r *http.Request,
	requestedFileOrDirectory string,
) {
	if containsRelativeBackPaths(r.URL.Path) {
		serveWaywardRequest(w, r)
		return
	}

	_, errFileInfo := os.Stat(requestedFileOrDirectory)
	if os.IsNotExist(errFileInfo) {
		serveWaywardRequest(w, r)
		return
	}

	http.ServeFile(w, r, requestedFileOrDirectory)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	requestedFileOrDirectory := webClientsDirectory + r.URL.Path
	serveStaticFiles(w, r, requestedFileOrDirectory)
}
