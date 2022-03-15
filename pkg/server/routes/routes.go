package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/gnanaprakash55/termus/pkg/roundrobin"
	"github.com/gnanaprakash55/termus/pkg/server/utils"
)

var mu sync.Mutex

var rr roundrobin.RoundRobin = utils.Init_roundrobin()

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s: %s%s", req.Method, req.Host, req.URL.Path)
	// var payload Payload

	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.Printf("Error reading body: %v", err)
	// 	panic(err)
	// }

	// req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// err = json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body))).Decode(&payload)
	// if err != nil {
	// 	log.Printf("Error parsng body: %v", err)
	// 	panic(err)
	// }

	//Get Proxy URL to redirect

	currentServer := rr.Next()

	fmt.Println(currentServer)

	if currentServer.GetIsDead() {

	}

	log.Printf("Redirecting to Proxy URL %s", currentServer.URL)

	var url *url.URL = currentServer.URL

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("%v is dead", url)
		currentServer.SetDead(true)
		HandleRequest(w, r)
	}
	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)

	// write message into kafka producer
	// ctx := context.Background()
	// message := "Redirecting to Proxy URL " + proxyURL
	// go kafka.Producer(ctx, message)

}
