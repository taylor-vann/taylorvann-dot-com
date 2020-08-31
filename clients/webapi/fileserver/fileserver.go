package fileserver

import (
	"net/http"
	"os"
	"io/ioutil"

	"log"
)

var (
	webClientsDirectory    = os.Getenv("WEB_CLIENTS_DIRECTORY")
	waywardRequestFilename = webClientsDirectory + "/lost/index.html"

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
	custom404, _ := ioutil.ReadFile(waywardRequestFilename)
	w.WriteHeader(http.StatusNotFound)
	w.Write(custom404)
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

	fileInfo, errFileInfo := os.Stat(requestedFileOrDirectory)
	if os.IsNotExist(errFileInfo) {
		serveWaywardRequest(w, r)
		return
	}

	if fileInfo.IsDir() {
		log.Println(requestedFileOrDirectory + "/index.html")
		_, errDirInfo := os.Stat(requestedFileOrDirectory + "/index.html")

		if os.IsNotExist(errDirInfo) {
			serveWaywardRequest(w, r)
			return
		}
	}

	http.ServeFile(w, r, requestedFileOrDirectory)
}

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, webClientsDirectory)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	requestedFileOrDirectory := webClientsDirectory + r.URL.Path
	serveStaticFiles(w, r, requestedFileOrDirectory)
}
