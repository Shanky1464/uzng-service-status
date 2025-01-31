package main

import (
	"fmt"

	"uniphore.com/uzng-service-status/internal/api/app"
	"uniphore.com/uzng-service-status/internal/handler"
	"uniphore.com/uzng-service-status/internal/handler/v1api"
	"uniphore.com/uzng-service-status/pkg/apm"
	"uniphore.com/uzng-service-status/pkg/lgr"
	"uniphore.com/uzng-service-status/pkg/metrics"
	"uniphore.com/uzng-service-status/pkg/router"
)

const (
	V1Path     = "/v1"
	HealthPath = "/health"
)

func main() {
	appConfig, err := app.NewConfig()
	if err != nil {
		lgr.Fatalf("Failed to initialize configuration: %v", err)
	}

	apm.Start()
	defer apm.Stop()

	lgr.Setup(appConfig.Logger)

	metrics, err := metrics.New(appConfig.Metrics)
	if err != nil {
		lgr.Fatalf("Failed to create DataDog metrics client: %s", err)
	}
	defer metrics.Close()

	router := router.New(appConfig.Router)

	helloWorldHandlerV1 := v1api.NewHelloWorld(metrics)

	// v1
	v1 := router.Group(V1Path)
	{
		v1.GET("/hello", helloWorldHandlerV1.Get)
	}

	// internal
	router.GET(HealthPath+"/liveness", handler.GetHealthLiveness)
	router.GET(HealthPath+"/readiness", handler.GetHealthReadiness)

	// serve
	router.Run(fmt.Sprintf(":%d", appConfig.Router.Port))
}
