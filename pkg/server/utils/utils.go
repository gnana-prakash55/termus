package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gnanaprakash55/termus/pkg/parsing"
)

const filename string = "termus.toml"

//Get Proxy URL by proxy condition
func GetProxyURL(proxyConditionRAW string) string {
	proxyCondition := strings.ToUpper(proxyConditionRAW)

	config := parsing.ParseConfig(filepath.Join("./config/", filename))

	if proxyCondition == "A" {
		return config.Servers[0]
	}

	if proxyCondition == "B" {
		return config.Servers[1]
	}

	if proxyCondition == "C" {
		return config.Servers[2]
	}

	return config.Default

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
