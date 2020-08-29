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
	hostname := r.Host
	redirectUrl := httpsDelimiter + hostname + r.RequestURI
	return redirectUrl
}

func PassToHttps(w http.ResponseWriter, r *http.Request) {
	redirectURL := getRedirectURL(r)
	http.Redirect(
		w,
		r,
		redirectURL,
		http.StatusMovedPermanently,
	)
}