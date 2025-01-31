//go:build unit

package apm_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/apm"
)

func TestAPMConfig(t *testing.T) {
	env := "env-value"
	service := "service-value"

	os.Setenv("DD_ENV", env)
	os.Setenv("DD_SERVICE", service)

	config, err := apm.NewConfig()

	assert.NoError(t, err, "Should not be error creating APM config")
	assert.Equal(t, env, config.Environment, "APM Datadog env should be equal")
	assert.Equal(t, service, config.Service, "APM Datadog service should be equal")

	os.Unsetenv("DD_ENV")
	os.Unsetenv("DD_SERVICE")
}

func TestAPMInvalidConfig(t *testing.T) {
	config, err := apm.NewConfig()

	assert.Error(t, err, "Should be error creating APM config")
	assert.Equal(t, apm.Config{}, config, "APM config should be empty")
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
}
