package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DB_URL"`
}

func Load() *Config {
	var cnf Config
	if err := envconfig.Process("", &cnf); err != nil {
		log.Fatal(err.Error())
	}
	return &cnf
}
