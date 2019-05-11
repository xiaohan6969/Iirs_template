package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

var Config *toml.Tree

func init() {
	path := "src/config/config.toml"
	config, err := toml.LoadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	Config = config
}
