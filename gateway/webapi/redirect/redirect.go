// brian taylor vann
// taylorvann dot com

package redirect

import (
	"net/http"
)

func getRedirectURL(r *http.Request) string {
	hostname := r.Host
	redirectUrl := "https://" + hostname + r.RequestURI
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