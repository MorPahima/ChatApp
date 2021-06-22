package cmd

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/palantir/stacktrace"
)

type Config struct {
	ServerHost string `env:"SERVER_HOST" envDefault:"localhost"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`
	MongoURI   string `env:"MONGO_URI"`
}

func (config Config) BaseUrl() string {
	return fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
}

func ParseConfig() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, stacktrace.Propagate(err, "")
	}

	return cfg, err
}
