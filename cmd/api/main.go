package main

import (
	"fmt"

	"uniphore.com/platform-hello-world-go/internal/api/app"
	"uniphore.com/platform-hello-world-go/internal/handler"
	"uniphore.com/platform-hello-world-go/internal/handler/v1api"
	"uniphore.com/platform-hello-world-go/pkg/apm"
	"uniphore.com/platform-hello-world-go/pkg/lgr"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
	"uniphore.com/platform-hello-world-go/pkg/router"
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
