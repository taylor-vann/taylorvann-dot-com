//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"
	"os"

	"fmt"

	"webapi/fileserver"
)

var (
	Environment = os.Getenv("STAGE")
)

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	if Environment == "DEVELOPMENT" {
		fmt.Println("dvelopment stuffs")
		mux.HandleFunc("/tests/", fileserver.ServeHomeFiles)
		mux.HandleFunc("/sign-in/tests/", fileserver.ServeSignInFiles)
		mux.HandleFunc("/internal/tests/", fileserver.ServeInternalFiles)
	}

	mux.HandleFunc("/", fileserver.ServeHomeApp)
	mux.HandleFunc("/scripts/", fileserver.ServeHomeFiles)
	mux.HandleFunc("/styles/", fileserver.ServeHomeFiles)

	mux.HandleFunc("/sign-in/", fileserver.ServeSignInFiles)

	mux.HandleFunc("/internal/", fileserver.ServeInternalApp)
	mux.HandleFunc("/internal/scripts/", fileserver.ServeInternalFiles)
	mux.HandleFunc("/internal/styles/", fileserver.ServeInternalFiles)
	
	return mux
}
