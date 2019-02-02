package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	host := "http://bar:9000"

	proxyH := newProxy(host)
	r.PathPrefix("/").Handler(proxyH)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func newProxy(host string) http.Handler {
	remote, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	return newReverseProxy(remote)
}

func newReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme

		req.URL.Host = target.Host
		req.Host = req.URL.Host

		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}

	return &httputil.ReverseProxy{Director: director}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
