package parsing

import (
	"log"

	"github.com/BurntSushi/toml"
)

// struct for parsing config file
type Config struct {
	Port    int
	Servers []string
	Default string
}

//To parse Configuration File
func ParseConfig(path string) Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
