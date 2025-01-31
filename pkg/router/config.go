package router

import (
	"github.com/caarlos0/env/v9"
	"uniphore.com/platform-hello-world-go/pkg/apm"
)

type Config struct {
	Mode string `env:"APP_MODE" envDefault:"debug"`
	Port int    `env:"APP_PORT" envDefault:"8080"`
	APM  apm.Config
}

func NewConfig() (Config, error) {
	var conf Config

	if err := env.Parse(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
