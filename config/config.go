package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	AppName  string `envconfig:"app_name"`
	AppQuote string `envconfig:"app_quote"`
	Version  string `envconfig:"version"`
	LogLevel string `envconfig:"log_level"`
}

// Key ...
var Key Config

func init() {
	err := envconfig.Process("jorah", &Key)
	if err != nil {
		log.Fatal(err.Error())
	}
}
