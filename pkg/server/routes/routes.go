package routes

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s: %s%s", req.Method, req.Host, req.URL.Path)
	fmt.Fprint(res, "Termus - A Reverse Proxy Server")
}
