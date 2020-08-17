package subdomains

import (
	"net/http"
)

type ProxyMux map[string]http.Handler

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

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname := r.Host
	subdomain := getSubdomain(hostname)
	mux := proxyMux[subdomain]

	if mux != nil {
		r.Header.Set("X-Forwarded-Proto", "https")
		mux.ServeHTTP(w, r)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}