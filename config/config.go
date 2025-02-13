package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort     string `envconfig:"APP_PORT" default:"8080"`
	AppLogLevel string `envconfig:"APP_LOG_LEVEL" default:"info"`

	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBDatabase string `envconfig:"DB_DATABASE" required:"true"`
	DBUsername string `envconfig:"DB_USERNAME" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBLogLevel string `envconfig:"DB_LOG_LEVEL" default:"info"`
}

var cfg *Config

func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg = &Config{}
	envconfig.MustProcess("", cfg)
	return cfg
}
