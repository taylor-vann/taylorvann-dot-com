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

		//	if no cookies redirect to"you made it to https" homepage / 404
		//		via dev.taylorvann.com/404 or taylorvann
		// mux.ServeHTTP(w, r)

		// for now get hostname
		json.NewEncoder(w).Encode(r.URL)
		return
	}

	// obscure print everything if hostname doesn't work
	json.NewEncoder(w).Encode(r.URL)

	http.Error(w, r.URL.Hostname(), 404)
}

func getRedirectURL(r *http.Request) string {
	return "https://" + r.Host + r.RequestURI
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

func CreateProxyMux() *ProxyMux {
	proxyMux := make(ProxyMux)
	for _, details := range *constants.RouteMap {
		url, errUrl := url.Parse(details.RequestedAddress)
		if errUrl != nil {
			continue
		}
		urlTarget, errUrlTarget := url.Parse(details.TargetAddress)
		if errUrlTarget != nil {
			continue
		}

		hostname := url.Hostname()

		proxyMux[hostname] = httputil.NewSingleHostReverseProxy(urlTarget)
	}

	return &proxyMux
}

func RedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", passToHttps)

	return mux
}