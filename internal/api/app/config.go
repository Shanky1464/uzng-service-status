package app

import (
	"uniphore.com/platform-hello-world-go/pkg/apm"
	"uniphore.com/platform-hello-world-go/pkg/lgr"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
	"uniphore.com/platform-hello-world-go/pkg/router"
)

type AppConfig struct {
	Logger  lgr.Config
	APM     apm.Config
	Metrics metrics.Config
	Router  router.Config
}

func NewConfig() (AppConfig, error) {
	logger, err := lgr.NewConfig()
	if err != nil {
		return AppConfig{}, err
	}

	apm, err := apm.NewConfig()
	if err != nil {
		return AppConfig{}, err
	}

	metrics, err := metrics.NewConfig()
	if err != nil {
		return AppConfig{}, err
	}

	router, err := router.NewConfig()
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Logger:  logger,
		APM:     apm,
		Metrics: metrics,
		Router:  router,
	}, nil
}
