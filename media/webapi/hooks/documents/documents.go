package documents

import (
	// "fmt"
	"fmt"
	"net/http"
	// "time"

	"webapi/hooks/constants"
	// "os"
)

type DocumentTokens struct {
	SessionToken      string
	CsrfToken         string
	DocumentToken     string
	DocumentCsrfToken string
}

// func createGuestSessionTokens(w http.ResponseWrite, r) {

// }

// func updateSessionTokens(w http.ResponseWrite, r) {
	
// }

func createOrExtendSessionTokens(w http.ResponseWriter, r *http.Request) {
	// get new guest document and guest session
	sessionToken, errSessionToken := r.Cookie(constants.SessionTokenCookie)
	csrfToken, errCsrfToken := r.Cookie(constants.CsrfTokenCookie)

	if sessionToken != nil &&
		csrfToken != nil &&
		errSessionToken == nil &&
		errCsrfToken == nil {
		// update tokens
		fmt.Println("got sessions to update")
		return
	}

	// get guest sessions
	fmt.Println("do new stuff")
}

// Homepage -
func Homepage(w http.ResponseWriter, r *http.Request) {
	createOrExtendSessionTokens(w, r)

	http.ServeFile(w, r, "/usr/local/web_clients/web_client/dist/index.html")
}

// Login -
func Login(w http.ResponseWriter, r *http.Request) {
	// if no cookie exists, fetch guest jwt and put as cookie
	// slap CSRF token into headers

	// if cookie exists, check against csrf,
	http.ServeFile(w, r, "/usr/local/web_clients/web_client_login/dist/index.html")
}

// Login -
func Logout(w http.ResponseWriter, r *http.Request) {
	// remove document and session cookies from header
	// if cookie exists, check against csrf,

	// intentionally put guest headers and return
	Homepage(w, r)
}
