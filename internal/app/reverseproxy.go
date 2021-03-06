package app

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// ReverseProxy basic structs to hold parameters
type ReverseProxy struct {
	IP          string `json:"ip"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Proxyport   string `json:"proxyPort"`
	HTTPTimeout int64  `json:"httpTimeout"`
	CertFile    string `json:"certfile"`
	KeyFile     string `json:"keyfile"`
}

// New returns a new initialization of the reverseproxy struct instance
func New(ip, port, protocol, proxyport, certfile, keyfile string, httptimeout int64) *ReverseProxy {
	return &ReverseProxy{
		IP:          ip,
		Port:        port,
		Protocol:    protocol,
		Proxyport:   proxyport,
		HTTPTimeout: httptimeout,
		CertFile:    certfile,
		KeyFile:     keyfile,
	}
}

// RunProxy runs the proxy
func (p *ReverseProxy) RunProxy() {
	urlStr := fmt.Sprintf("%s://%s:%s", p.Protocol, p.IP, p.Port)
	remote, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to parse the url: %v due to %v", urlStr, err))
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	tr := &http.Transport{IdleConnTimeout: time.Duration(p.HTTPTimeout) * time.Second}

	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	proxy.Transport = tr
	http.HandleFunc("/", handler(proxy))
	log.Printf("Serving reverse proxy on %v", p.Proxyport)
	proxyURL := fmt.Sprintf("0.0.0.0:%v", p.Proxyport)
	err = http.ListenAndServeTLS(proxyURL, p.CertFile, p.KeyFile, nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to start the tls proxy due to %v", err))
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		p.ServeHTTP(w, r)
	}
}
