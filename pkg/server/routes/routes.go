package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gnanaprakash55/termus/pkg/parsing"
)

const filename string = "termus.toml"

type Payload struct {
	ProxyCondition string `json:"proxy_condition"`
}

func getProxyURL(proxyConditionRAW string) string {
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

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
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

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s: %s%s", req.Method, req.Host, req.URL.Path)
	var payload Payload

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body))).Decode(&payload)
	if err != nil {
		log.Printf("Error parsng body: %v", err)
		panic(err)
	}

	proxyURL := getProxyURL(payload.ProxyCondition)

	serveReverseProxy(proxyURL, res, req)

}
