package apm

import "github.com/caarlos0/env/v9"

type Config struct {
	Environment string `env:"DD_ENV,required"`
	Service     string `env:"DD_SERVICE,required"`
}

func NewConfig() (Config, error) {
	var conf Config

	if err := env.Parse(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
