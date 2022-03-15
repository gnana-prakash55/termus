package parsing

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// struct for parsing config file
type Config struct {
	Port     int
	Backends []Backend
}

type Backend struct {
	URL string
}

const filename = "termus.toml"

//To parse Configuration File
func ParseConfig() Config {
	var config Config
	if _, err := toml.DecodeFile(filepath.Join("./config/", filename), &config); err != nil {
		log.Fatal(err)
	}
	return config
}

//Get Port from Config File
func GetPort() int {
	var config Config
	if _, err := toml.DecodeFile(filepath.Join("./config/", filename), &config); err != nil {
		log.Fatal(err)
	}
	return config.Port
}

//Get Servers from Config File
func GetServers() []Backend {
	var config Config
	if _, err := toml.DecodeFile(filepath.Join("./config/", filename), &config); err != nil {
		log.Fatal(err)
	}
	return config.Backends
}
