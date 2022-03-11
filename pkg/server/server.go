package server

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gnanaprakash55/termus/pkg/parsing"
	"github.com/gnanaprakash55/termus/pkg/server/routes"
)

const filename string = "termus.toml"

func Start() {
	var config parsing.Config
	config = parsing.ParseConfig(filepath.Join("./config/", filename))

	http.HandleFunc("/", routes.HandleRequest)

	log.Printf("Starting Server at PORT %s", strconv.Itoa(config.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))

}
