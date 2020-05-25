package subdomains

import (
	"net/http"

	// "webapi/server/routes/constants"
)

const XForwardedProto = "X-Forwarded-Proto"
const https = "https"

type ProxyMux map[string]http.Handler

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname := r.Host
	subdomain := getSubdomain(hostname)
	mux := proxyMux[subdomain]

	if mux != nil {
		// add header here
		r.Header.Set("X-Forwarded-Proto", "https")
		mux.ServeHTTP(w, r)
		return
	}

	// reroute to 404
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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