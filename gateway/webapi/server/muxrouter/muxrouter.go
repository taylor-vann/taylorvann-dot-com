//	brian taylor vann
//	briantaylorvann dot com

package muxrouter

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"webapi/server/muxrouter/redirect"
	"webapi/server/muxrouter/subdomains"
)

var (
	Enviornment = os.Getenv("STAGE")

	AuthnAddress = os.Getenv("AUTHN_URL_ADDRESS")
	ClientsAddress = os.Getenv("CLIENTS_URL_ADDRESS")
	LogsAddress = os.Getenv("LOGS_URL_ADDRESS")
	MailAddress = os.Getenv("MAIL_URL_ADDRESS")
	MediaAddress = os.Getenv("MEDIA_URL_ADDRESS")

	RouteMap = map[string]string{
		"briantaylorvann": ClientsAddress,
		"authn": AuthnAddress,
		"logs": LogsAddress,
		"mail": MailAddress,
		"media": MediaAddress,
	}
)

func CreateProxyMux() *subdomains.ProxyMux {
	proxyMux := make(subdomains.ProxyMux)
	for subdomain, address := range RouteMap {
		url, errUrl := url.Parse(address)
		if errUrl != nil {
			continue
		}
		proxyMux[subdomain] = httputil.NewSingleHostReverseProxy(url)
	}

	return &proxyMux
}

func CreateRedirectToHttpsMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirect.PassToHttps)

	return mux
}
