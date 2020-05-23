package subdomains

import (
	"encoding/json"
	"net/http"
)

type ProxyMux map[string]http.Handler

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname := r.Host
	subdomain := getSubdomain(hostname)
	mux := proxyMux[subdomain]

	if mux != nil {
		mux.ServeHTTP(w, r)
		return
	}

	// reroute to 404
	mux = proxyMux[constants.BrianTaylorVann]
	mux.ServeHTTP(w, r)
}

func getSubdomain(hostname string) string {
	index, subdomain := getDelimiterIndex(0, hostname, ".")
	if subdomain == "www" {
		index += 1
		index, subdomain = getDelimiterIndex(index, hostname, ".")
	}
	return subdomain
}

func getDelimiterIndex(index int, hostname string, delimiter string) (int, string) {
	byteDelimiter := byte(delimiter[0])
	initialIndex := index
	searchIndex := initialIndex + 1

	for searchIndex < len(hostname) {
		if hostname[searchIndex] == byteDelimiter {
			break
		}
		searchIndex += 1
	}

	return searchIndex, hostname[initialIndex:searchIndex]
}