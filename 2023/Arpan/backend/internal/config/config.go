package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port string `envconfig:"PORT" default:"8080" required:"true" split_words:"true"`
}

func LoadAppConfig() *AppConfig {
	c := &AppConfig{}
	if err := envconfig.Process("", c); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	return c
}
