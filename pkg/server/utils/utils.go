package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/roundrobin"
)

var urls []*url.URL

var rr roundrobin.RoundRobin = init_roundrobin()

//initialize round robin algo
func init_roundrobin() roundrobin.RoundRobin {

	var servers = parsing.GetServers()

	for i := 0; i < len(servers); i++ {
		url, _ := url.Parse(servers[i])
		urls = append(urls, url)
	}

	rr, err := roundrobin.New(urls)
	if err != nil {
		panic(err)
	}

	return rr
}

//Health Check for servers
func healthCheck(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

//Get Proxy URL by doing Health Check
func GetProxyURL() string {
	var url string = rr.Next().String()
	if !healthCheck(url) {
		return GetProxyURL()
	}
	return url
}

//Serving as Reverse Proxy
func ServeReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
