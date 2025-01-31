package lgr

import "github.com/caarlos0/env/v9"

type Config struct {
	Level       string `env:"LOG_LEVEL" envDefault:"INFO"`
	TraceCaller bool   `env:"LOG_TRACE_CALLER" envDefault:"true"`
}

func NewConfig() (Config, error) {
	var conf Config

	if err := env.Parse(&conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
