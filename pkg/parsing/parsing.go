package parsing

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port    int
	Servers []string
	Default string
}

func ParseConfig(path string) Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
