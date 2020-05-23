//	brian taylor vann
//	briantaylorvann dot com

package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"webapi/redirect"
	"webapi/routes/constants"
	"webapi/routes/subdomains"
)

func CreateProxyMux() *subdomains.ProxyMux {
	proxyMux := make(subdomains.ProxyMux)
	for subdomain, address := range *constants.Routes {
		url, errUrl := url.Parse(address)
		if errUrl != nil {
			continue
		}
		proxyMux[subdomain] = httputil.NewSingleHostReverseProxy(url)
	}

	return &proxyMux
}


func RedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirect.PassToHttps)

	return mux
}
