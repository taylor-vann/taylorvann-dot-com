// brian taylor vann
// taylorvann dot com

package redirect

import (
	"net/http"
)

func getRedirectURL(r *http.Request) string {
	hostname := r.URL.Hostname()
	return "https://" + hostname + r.RequestURI
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