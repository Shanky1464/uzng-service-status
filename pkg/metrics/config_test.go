//go:build unit

package metrics_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
)

func TestMetricsConfig(t *testing.T) {
	agentHost := "agent-host-value"
	os.Setenv("DD_AGENT_HOST", agentHost)

	config, err := metrics.NewConfig()

	assert.NoError(t, err, "Should not be error creating metrics config")
	assert.Equal(t, agentHost, config.Host, "Datadog agent host should be equal")

	os.Unsetenv("DD_AGENT_HOST")
}

func TestMetricsInvalidConfig(t *testing.T) {
	config, err := metrics.NewConfig()

	assert.Error(t, err, "Should be error creating log config")
	assert.Equal(t, metrics.Config{}, config, "Metrics config should be empty")
	assert.Contains(
		t,
		err.Error(),
		`required environment variable "DD_AGENT_HOST" is not set`,
		"Datadog Metrics agent host should be set",
	)
}
