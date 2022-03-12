package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gnanaprakash55/termus/pkg/kafka"
	"github.com/gnanaprakash55/termus/pkg/server/utils"
)

//Request Payload
type Payload struct {
	ProxyCondition string `json:"proxy_condition"`
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

	proxyURL := utils.GetProxyURL(payload.ProxyCondition)

	log.Printf("Proxy Condition %s, Redirecting to Proxy URL %s", payload.ProxyCondition, proxyURL)

	// write message into kafka producer
	ctx := context.Background()
	message := "Proxy Condition " + payload.ProxyCondition + ", Redirecting to Proxy URL " + proxyURL
	kafka.Producer(ctx, message)

	utils.ServeReverseProxy(proxyURL, res, req)

}
