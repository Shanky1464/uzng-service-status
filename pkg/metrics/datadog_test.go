//go:build unit

package metrics_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
)

func TestMetrics(t *testing.T) {
	m, err := metrics.New(metrics.Config{Host: "localhost"})

	assert.NoError(t, err, "Should not be error creating metircs agent")
	assert.NotNil(t, m)
}
