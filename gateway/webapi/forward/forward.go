// brian taylor vann
// briataylorvann

package forward

import (
	"net/http"
)

type ProxyMux map[string]http.Handler

const (
	httpsStatement     = "https"
	XForwardedFor    = "X-Forwarded-For"
)

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	host := r.URL.Host
	// r.URL.Scheme = "https"
	// scheme://host/path/

	mux := proxyMux[host]

	// u, err := url.Parse(r.RequestURI)
	// u.Scheme = "https"

	if mux != nil {
		r.Header.Set(XForwardedFor, r.RemoteAddr)
		mux.ServeHTTP(w, r)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
