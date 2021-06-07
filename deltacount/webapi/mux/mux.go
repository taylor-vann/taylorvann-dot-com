//	brian taylor vann
//	briantaylorvann dot com

package muxrouter

import (
	"net/http"
	"net/url"
	// import limiter x
)

const (
	homeRoute      = "/"
	httpsScheme    = "https"
	XForwardedFor  = "X-Forwarded-For"
	XForwardedHost = "X-Forwarded-Host"
	emptyString    = ""
)

// json response
type LimiterRequestPayload struct {
	Address	string	`json:"address"`
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	// take user details,

	// take details
	// return quickness

	dest, errDest := url.Parse(r.URL.String())
	if errDest != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	dest.Scheme = httpsScheme
	destStr := dest.String()

	http.Redirect(
		w,
		r,
		destStr,
		http.StatusMovedPermanently,
	)
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, redirectToHTTPS)

	return mux
}
