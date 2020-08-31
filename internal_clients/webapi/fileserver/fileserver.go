package fileserver

import (
	"net/http"
	"os"

	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/sessionx"
	"github.com/taylor-vann/weblog/toolbox/golang/infraclientx/verifyx"
)

type ValidateParams = verifyx.HasRoleFromSessionParams

const (
	SessionCookieHeader = "briantaylorvann.com_session"
	InternalAdmin       = "INTERNAL_ADMIN"
)

var (
	Environment = os.Getenv("STAGE")

	webClientsDirectory    = os.Getenv("WEB_CLIENTS_DIRECTORY")
	waywardRequestFilename = webClientsDirectory + "lost/index.html"
	relativeRune           = []byte(".")[0]
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

func validateInternalUser(
	w http.ResponseWriter,
	p *ValidateParams,
) bool {
	isValidSession := verifyx.IsSessionValid(&verifyx.IsSessionValidParams{
		Environment:        p.Environment,
		InfraSessionCookie: p.InfraSessionCookie,
		SessionCookie:      p.SessionCookie,
	})
	if !isValidSession {
		return false
	}

	if verifyx.HasRoleFromSession(p) {
		return true
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
		_, errDirInfo := os.Stat(requestedFileOrDirectory + "/index.html")

		if os.IsNotExist(errDirInfo) {
			serveWaywardRequest(w, r)
			return
		}
	}

	http.ServeFile(w, r, requestedFileOrDirectory)
}

func ServeSignIn(w http.ResponseWriter, r *http.Request) {
	requestedFileOrDirectory := webClientsDirectory + r.URL.Path
	serveStaticFiles(w, r, requestedFileOrDirectory)
}

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, webClientsDirectory)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	requestedFileOrDirectory := webClientsDirectory + r.URL.Path

	sessionCookie, _ := r.Cookie(SessionCookieHeader)

	isValidInternalUser := validateInternalUser(w, &ValidateParams{
		Environment:        Environment,
		InfraSessionCookie: sessionx.InfraSession,
		SessionCookie:      sessionCookie,
		Organization:       InternalAdmin,
	},
	)
	if !isValidInternalUser {
		return
	}

	serveStaticFiles(w, r, requestedFileOrDirectory)
}
