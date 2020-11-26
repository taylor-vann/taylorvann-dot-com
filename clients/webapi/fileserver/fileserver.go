package fileserver

import (
	"io/ioutil"
	"net/http"
	"os"

	"fmt"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/verifyx"
)

type ValidateParams = verifyx.HasRoleFromSessionParams

const (
	SessionCookieHeader = "briantaylorvann.com_session"
	InternalAdmin       = "INTERNAL_ADMIN"
	
	IndexHTML = "/index.html"
)

var (
	Environment = os.Getenv("STAGE")
	webClientsDirectory = os.Getenv("WEB_CLIENTS_DIRECTORY")

	unauthorizedFilename = webClientsDirectory + "/unauthorized/unauthorized/index.html"
	waywardFilename = webClientsDirectory + "/lost/lost/index.html"
	signInDirectory = webClientsDirectory + "/sign-in/"
	homeResourcesDirectory = webClientsDirectory + "/home/home/"
	homeFilename = webClientsDirectory + "/home/home/index.html"
	internalResourcesDirectory = webClientsDirectory + "/internal/"
	internalFilename = webClientsDirectory + "/internal/internal/index.html"

	relativeRune = []byte(".")[0]
)

var (
	custom401, errCustom401 = ioutil.ReadFile(unauthorizedFilename)
	custom404, errCustom404 = ioutil.ReadFile(waywardFilename)
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

func isValidInternalUser(r *http.Request) bool {
	sessionCookie, _ := r.Cookie(SessionCookieHeader)
	if sessionCookie == nil {
		return false
	}

	if !verifyx.IsSessionValid(&verifyx.IsSessionValidParams{
		Environment:        Environment,
		InfraSessionCookie: sessionx.InfraSession,
		SessionCookie:      sessionCookie,
	}) {
		return false
	}

	if verifyx.HasRoleFromSession(&ValidateParams{
		Environment:        Environment,
		InfraSessionCookie: sessionx.InfraSession,
		SessionCookie:      sessionCookie,
		Organization:       InternalAdmin,
	}) {
		return true
	}

	return false
}

func serveUnauthorizedRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(custom401)
}

func serveWaywardRequest(w http.ResponseWriter, r *http.Request) {
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
		_, errDirInfo := os.Stat(requestedFileOrDirectory + IndexHTML)
		if os.IsNotExist(errDirInfo) {
			serveWaywardRequest(w, r)
			return
		}
	}

	http.ServeFile(w, r, requestedFileOrDirectory)
}

func ServeInternalFiles(w http.ResponseWriter, r *http.Request) {
	if !isValidInternalUser(r) {
		serveUnauthorizedRequest(w, r)
		return
	}

	serveStaticFiles(w, r, internalResourcesDirectory + r.URL.Path)
}

func ServeInternalApp(w http.ResponseWriter, r *http.Request) {
	if !isValidInternalUser(r) {
		serveUnauthorizedRequest(w, r)
		return
	}

	serveStaticFiles(w, r, internalFilename)
}

func ServeSignInFiles(w http.ResponseWriter, r *http.Request) {
	serveStaticFiles(w, r, signInDirectory + r.URL.Path)
}

func ServeHomeFiles(w http.ResponseWriter, r *http.Request) {
	serveStaticFiles(w, r, homeResourcesDirectory + r.URL.Path)
}

func ServeHomeApp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, homeFilename)
}