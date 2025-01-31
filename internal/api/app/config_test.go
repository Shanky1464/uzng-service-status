//go:build unit

package app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/internal/api/app"
)

func TestAppConfig(t *testing.T) {
	logLevel := "log-level-value"
	logTraceCaller := "false"
	apmEnv := "apm-env-value"
	apmService := "apm-service-value"
	metricsAgentHost := "metrics-host-value"
	appMode := "app-mode-value"
	appPort := "1234"

	os.Setenv("LOG_LEVEL", logLevel)
	os.Setenv("LOG_TRACE_CALLER", logTraceCaller)
	os.Setenv("DD_ENV", apmEnv)
	os.Setenv("DD_SERVICE", apmService)
	os.Setenv("DD_AGENT_HOST", metricsAgentHost)
	os.Setenv("APP_MODE", appMode)
	os.Setenv("APP_PORT", appPort)

	config, err := app.NewConfig()

	assert.NoError(t, err, "Should not be error creating router config")
	// Logger
	assert.Equal(t, logLevel, config.Logger.Level, "Logger level should be equal")
	assert.False(t, config.Logger.TraceCaller, "Logger trace caller should be false")
	// APM
	assert.Equal(t, apmEnv, config.APM.Environment, "APM Datadog env should be equal")
	assert.Equal(t, apmService, config.APM.Service, "APM Datadog service should be equal")
	// Metrics
	assert.Equal(t, metricsAgentHost, config.Metrics.Host, "Metrics Datadog agent host should be equal")
	// Router
	assert.Equal(t, appMode, config.Router.Mode, "App mode should be equal")
	assert.Equal(t, appPort, fmt.Sprint(config.Router.Port), "App port should be equal")
	assert.Equal(t, apmEnv, config.Router.APM.Environment, "APM Datadog environment should be equal")
	assert.Equal(t, apmService, config.Router.APM.Service, "APM Datadog service should be equal")

	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("LOG_TRACE_CALLER")
	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
	os.Unsetenv("DD_AGENT_HOST")
	os.Unsetenv("APP_MODE")
	os.Unsetenv("APP_PORT")
}

func TestAppLogInvalidConfig(t *testing.T) {
	// Prerequisites
	apmEnv := "apm-env-value"
	apmService := "apm-service-value"
	metricsAgentHost := "metrics-host-value"
	os.Setenv("DD_ENV", apmEnv)
	os.Setenv("DD_SERVICE", apmService)
	os.Setenv("DD_AGENT_HOST", metricsAgentHost)

	invalidValue := "invalid-value"
	os.Setenv("LOG_TRACE_CALLER", invalidValue)

	config, err := app.NewConfig()

	assert.Error(t, err, "Should be error creating app config")
	assert.Equal(t, app.AppConfig{}, config, "App config should be empty")
	assert.Contains(
		t,
		err.Error(),
		fmt.Sprintf(
			`parse error on field "TraceCaller" of type "bool": strconv.ParseBool: parsing "%s": invalid syntax`,
			invalidValue),
		"Logger trace caller value should be bool type",
	)

	os.Unsetenv("LOG_TRACE_CALLER")
	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
	os.Unsetenv("DD_AGENT_HOST")
}

func TestAppAPMConfigError(t *testing.T) {
	// Prerequisites
	metricsAgentHost := "metrics-host-value"
	os.Setenv("DD_AGENT_HOST", metricsAgentHost)

	config, err := app.NewConfig()

	assert.Error(t, err, "Should be error creating app config")
	assert.Equal(t, app.AppConfig{}, config, "App config should be empty")
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

	os.Unsetenv("DD_AGENT_HOST")
}

func TestAppMetricsConfigError(t *testing.T) {
	// Prerequisites
	apmEnv := "apm-env-value"
	apmService := "apm-service-value"
	os.Setenv("DD_ENV", apmEnv)
	os.Setenv("DD_SERVICE", apmService)

	config, err := app.NewConfig()

	assert.Error(t, err, "Should be error creating app config")
	assert.Equal(t, app.AppConfig{}, config, "App config should be empty")
	assert.Contains(
		t,
		err.Error(),
		`required environment variable "DD_AGENT_HOST" is not set`,
		"Metrics Datadog agent host env variable should be set",
	)

	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
}

func TestAppRouterInvalidConfig(t *testing.T) {
	// Prerequisites
	apmEnv := "apm-env-value"
	apmService := "apm-service-value"
	metricsAgentHost := "metrics-host-value"
	os.Setenv("DD_ENV", apmEnv)
	os.Setenv("DD_SERVICE", apmService)
	os.Setenv("DD_AGENT_HOST", metricsAgentHost)

	invalidValue := "invalid-value"
	os.Setenv("APP_PORT", invalidValue)

	config, err := app.NewConfig()

	assert.Error(t, err, "Should be error creating app config")
	assert.Equal(t, app.AppConfig{}, config, "App config should be empty")
	assert.Contains(
		t,
		err.Error(),
		fmt.Sprintf(
			`parse error on field "Port" of type "int": strconv.ParseInt: parsing "%s": invalid syntax`,
			invalidValue),
		"App port value should be int type",
	)

	os.Unsetenv("APP_PORT")
	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
	os.Unsetenv("DD_AGENT_HOST")
}
