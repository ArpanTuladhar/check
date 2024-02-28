package config

import (
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type Env string

const (
	EnvProduction Env = "production"
	EnvStaging    Env = "staging"
	EnvLocal      Env = "local"
)

func ConstructDSN(conf *Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s",
		conf.DBUser,
		conf.DBPass,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		url.QueryEscape(conf.DBLoc),
	)
}

type Config struct {
	Env        Env    `required:"true" default:"local"`
	ServerPort int    `default:"8080" required:"true" split_words:"true"`
	DBHost     string `default:"127.0.0.1" required:"true" split_words:"true"`
	DBPort     int    `default:"3307" required:"true" split_words:"true"`
	DBUser     string `default:"root" required:"true" split_words:"true"`
	DBPass     string `default:"password" required:"true" split_words:"true"`
	DBName     string `default:"todo_dev" required:"true" split_words:"true"`
	DBLoc      string `default:"Asia/Tokyo" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &c, nil
}
