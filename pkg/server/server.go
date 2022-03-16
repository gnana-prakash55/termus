package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/server/routes"
	"github.com/gnanaprakash55/termus/pkg/server/utils"
)

// To start Server
func Start() {
	var PORT int
	PORT = parsing.GetPort()

	go utils.HealthCheck()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(PORT),
		Handler: http.HandlerFunc(routes.HandleRequest),
	}

	log.Printf("Starting Server at PORT %s", strconv.Itoa(PORT))
	log.Fatal(server.ListenAndServe())

}
