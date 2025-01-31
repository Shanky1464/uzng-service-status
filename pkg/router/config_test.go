//go:build unit

package router_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/router"
)

func TestRouterConfig(t *testing.T) {
	mode := "mode-value"
	port := "1234"
	apmEnv := "apm-env-value"
	apmService := "apm-service-value"

	os.Setenv("APP_MODE", mode)
	os.Setenv("APP_PORT", port)
	os.Setenv("DD_ENV", apmEnv)
	os.Setenv("DD_SERVICE", apmService)

	config, err := router.NewConfig()

	assert.NoError(t, err, "Should not be error creating router config")
	assert.Equal(t, mode, config.Mode, "App mode should be equal")
	assert.Equal(t, port, fmt.Sprint(config.Port), "App port should be equal")
	assert.Equal(t, apmEnv, config.APM.Environment, "APM environment should be equal")
	assert.Equal(t, apmService, config.APM.Service, "APM service should be equal")

	os.Unsetenv("APP_MODE")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
}

func TestRouterInvalidConfig(t *testing.T) {
	port := "invalid-value"
	os.Setenv("APP_PORT", port)

	config, err := router.NewConfig()

	assert.Error(t, err, "Should be error creating log config")
	assert.Equal(t, router.Config{}, config, "Router config should be empty")
	assert.Contains(
		t,
		err.Error(),
		fmt.Sprintf(
			`parse error on field "Port" of type "int": strconv.ParseInt: parsing "%s": invalid syntax`,
			port),
		"App port value should be int type",
	)
	assert.Contains(
		t,
		err.Error(),
		`required environment variable "DD_ENV" is not set`,
		"APM Datadog environment env variable should be set",
	)
	assert.Contains(
		t,
		err.Error(),
		`required environment variable "DD_SERVICE" is not set`,
		"APM Datadog service env variable should be set",
	)

	os.Unsetenv("APP_PORT")
}
