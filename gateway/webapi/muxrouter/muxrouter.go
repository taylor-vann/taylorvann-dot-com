//	brian taylor vann
//	briantaylorvann dot com

package muxrouter

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"webapi/details"
	"webapi/redirect"
	"webapi/forward"
)

var (
	homeRoute = "/"
)

func CreateProxyMux() *forward.ProxyMux {
	proxyMux := make(forward.ProxyMux)

	for dest, target := range details.Details.Routes {
		url, errUrl := url.Parse(target)
		if errUrl != nil {
			continue
		}
		proxyMux[dest] = httputil.NewSingleHostReverseProxy(url)
	}

	return &proxyMux
}

func CreateRedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, redirect.RedirectToHTTPS)

	return mux
}
