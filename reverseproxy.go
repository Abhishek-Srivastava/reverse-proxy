package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	HTTPTimeout = 10
)

func main() {
	remote, err := url.Parse("https://172.16.34.100")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	tr := &http.Transport{IdleConnTimeout: HTTPTimeout * time.Second}

	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	proxy.Transport = tr
	http.HandleFunc("/", handler(proxy))
	log.Println("Serving rproxy on 8080")
	err = http.ListenAndServeTLS("0.0.0.0:8080", "revpro.crt", "revpro.key", nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
