//	brian taylor vann
//	taylorvann dot com

//  Routes Gateway
//	Keep routes separate and isolated for easier scaling.
//	Each route could potentially be replaced by a simple
//	http request to an external service.

// Package routes -
package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
	"webapi/routes/constants"
)

type ProxyMux map[string]http.Handler

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := proxyMux[r.URL.Hostname()]
	if mux != nil {
		// check for document cookies
		//	if present, send to authn for verification
		//		if valid, continue onto request

		//	if no cookies redirect to homepage / 404
		//		via dev.taylorvann.com/404 or taylorvann
		// mux.ServeHTTP(w, r)
		json.NewEncoder(w).Encode("you made it to https")
		return
	}
	json.NewEncoder(w).Encode(r.URL)

	http.Error(w, r.URL.Hostname(), 404)
}

func getRedirectURL(r *http.Request) string {
	return "https://" + r.URL.Hostname() + ": 3005" + r.URL.String()
}

func passToHttps(w http.ResponseWriter, r *http.Request) {
	redirectURL := getRedirectURL(r)

	http.Redirect(
		w,
		r,
		redirectURL,
		http.StatusMovedPermanently,
	)
}

func helloHttps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("we're in https! " + r.URL.Hostname())
}

func CreateProxyMux() *ProxyMux {
	// mux := http.NewServeMux()

	// mux.HandleFunc("/", helloHttps)

	// return mux
	proxyMux := make(ProxyMux)
	for _, details := range *constants.RouteMap {
		url, errUrl := url.Parse(details.RequestedAddress)
		if errUrl != nil {
			continue
		}

		hostname := url.Hostname()

		urlTarget, errUrlTarget := url.Parse(details.TargetAddress)
		if errUrlTarget != nil {
			continue
		}

		proxyMux[hostname] = httputil.NewSingleHostReverseProxy(urlTarget)
	}

	return &proxyMux
}

func RedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", passToHttps)

	return mux
}
