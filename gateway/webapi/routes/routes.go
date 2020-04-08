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
	subdomain := getSubdomain(r.Host)
	mux := proxyMux[subdomain]

	if mux != nil {
		// check for document cookies
		//	if present, send to authn for verification
		//		if valid, continue onto request

		//	if no cookies redirect to"you made it to https" homepage / 404
		//		via dev.taylorvann.com/404 or taylorvann
		// mux.ServeHTTP(w, r)

		// for now get hostname
		json.NewEncoder(w).Encode("subdomain! " + r.Host)
		return
	}

	json.NewEncoder(w).Encode(r.Host)
}

func getSubdomain(host string) string {
	initialIndex := 0
	delimiterIndex := getDelimiterIndex(initialIndex, host, ".")

	subdomain := host[initialIndex:delimiterIndex]
	if subdomain == "www" {
		initialIndex = delimiterIndex + 1
		delimiterIndex = getDelimiterIndex(initialIndex, host, ".")
		subdomain = host[initialIndex:delimiterIndex]
	}

	return subdomain
}

func getDelimiterIndex(index int, host string, delimiter string) int {
	byteDelimiter := byte(delimiter[0])
	initialIndex := index
	searchIndex := initialIndex + 1

	for searchIndex < len(host) {
		if host[searchIndex] == byteDelimiter {
			break
		}
		searchIndex += 1
	}

	return searchIndex
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
		urlTarget, errUrlTarget := url.Parse(details.TargetAddress)
		if errUrlTarget != nil {
			continue
		}
		proxyMux[details.SubDomain] = httputil.NewSingleHostReverseProxy(urlTarget)
	}

	return &proxyMux
}

func RedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", passToHttps)

	return mux
}
