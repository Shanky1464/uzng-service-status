package app

import (
	"uniphore.com/uzng-service-status/pkg/apm"
	"uniphore.com/uzng-service-status/pkg/lgr"
	"uniphore.com/uzng-service-status/pkg/metrics"
	"uniphore.com/uzng-service-status/pkg/router"
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
