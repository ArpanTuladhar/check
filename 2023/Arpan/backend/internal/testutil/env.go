package testutil

import (
	"github.com/kelseyhightower/envconfig"
)

type TestEnv struct {
	DBHost string `default:"127.0.0.1" envconfig:"DB_HOST"`
	DBPort int    `default:"3307" envconfig:"DB_PORT"`
	DBUser string `default:"root" envconfig:"DB_USER"`
	DBPass string `default:"password" envconfig:"DB_PASS"`
	DBName string `default:"todo_test" envconfig:"DB_NAME"`
}

func LoadEnv() *TestEnv {
	var env TestEnv
	if err := envconfig.Process("", &env); err != nil {
		panic(err)
	}
	return &env
}
