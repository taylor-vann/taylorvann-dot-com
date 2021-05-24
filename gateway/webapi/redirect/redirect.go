// brian taylor vann
// briantaylorvann dot com

package redirect

import (
	"net/http"
)

const (
	httpsDelimiter = "https://"
)

func getRedirectURL(r *http.Request) string {
	host := r.Host
	redirectUrl := httpsDelimiter + host + r.RequestURI
	return redirectUrl
}

func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	redirectURL := getRedirectURL(r)
	http.Redirect(
		w,
		r,
		redirectURL,
		http.StatusMovedPermanently,
	)
}
