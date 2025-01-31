package metrics

import "github.com/caarlos0/env/v9"

type Config struct {
	Host string `env:"DD_AGENT_HOST,required"`
}

func NewConfig() (Config, error) {
	var conf Config

	if err := env.Parse(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
