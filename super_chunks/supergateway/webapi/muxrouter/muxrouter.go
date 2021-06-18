//	brian taylor vann
//	briantaylorvann dot com

package muxrouter

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	homeRoute      = "/"
	httpsScheme    = "https"
	XForwardedFor  = "X-Forwarded-For"
	XForwardedHost = "X-Forwarded-Host"
	emptyString    = ""
)

type ProxyMux map[string]http.Handler

func RedactURL(fullURL *url.URL, err error) (*url.URL, error) {
	if err != nil {
		return nil, err
	}

	redactedURL, errRedactedURL := url.Parse(fullURL.String())
	redactedURL.RawQuery = emptyString
	redactedURL.Fragment = emptyString

	return redactedURL, errRedactedURL
}

func RedactURLFromString(fullURL string, err error) (*url.URL, error) {
	redactedURL, errRedactedURL := url.Parse(fullURL)
	return RedactURL(redactedURL, errRedactedURL)
}

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	redactedURL, errRedactedURL := RedactURL(r.URL, nil)
	if errRedactedURL != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	proxyKey := redactedURL.String()
	mux, muxFound := proxyMux[proxyKey]
	if muxFound == false {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	r.Header.Set(XForwardedFor, r.RemoteAddr)
	r.Header.Set(XForwardedHost, r.Host)

	mux.ServeHTTP(w, r)
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	dest, errDest := url.Parse(r.URL.String())
	if errDest != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	dest.Scheme = httpsScheme
	destStr := dest.String()

	http.Redirect(
		w,
		r,
		destStr,
		http.StatusMovedPermanently,
	)
}

func CreateHTTPSMux(routes *map[string]string) (*ProxyMux, error) {
	proxyMux := make(ProxyMux)

	for dest, target := range *routes {
		destURL, errDestURL := RedactURLFromString(dest, nil)
		if errDestURL != nil {
			return nil, errDestURL
		}

		targetURL, errTargetURL := RedactURLFromString(target, nil)
		if errTargetURL != nil {
			return nil, errTargetURL
		}

		destRedacted := destURL.String()
		proxyMux[destRedacted] = httputil.NewSingleHostReverseProxy(targetURL)
	}

	return &proxyMux, nil
}

func CreateRedirectMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, redirectToHTTPS)

	return mux
}
