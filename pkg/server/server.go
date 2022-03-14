package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/server/routes"
)

// To start Server
func Start() {
	var PORT int
	PORT = parsing.GetPort()

	http.HandleFunc("/", routes.HandleRequest)

	log.Printf("Starting Server at PORT %s", strconv.Itoa(PORT))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))

}
