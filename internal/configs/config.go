// Package configs provides a struct for managing .env variables
package configs

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	*HTTPConfig
}

type HTTPConfig struct {
	Port int `env:"PORT"`
}

func LoadConfig() *Config {
	godotenv.Load(".env")

	conf := &Config{}
	errors := ParseConfig(conf)

	if len(errors) > 0 {
		printErrors(errors)
	}

	return conf
}

func printErrors(errs []error) {
	var msg strings.Builder
	msg.WriteString("config errors:\n")
	for _, err := range errs {
		fmt.Fprintf(&msg, "  - %s\n", err.Error())
	}
	log.Fatal(msg.String())
}
