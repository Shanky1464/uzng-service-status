//go:build unit

package lgr_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/lgr"
)

func TestLoggerConfig(t *testing.T) {
	level := "level-value"
	os.Setenv("LOG_LEVEL", level)

	config, err := lgr.NewConfig()

	assert.NoError(t, err, "Should not be error creating logger config")
	assert.Equal(t, level, config.Level, "Log level should be equal")
}

func TestLoggerInvalidConfig(t *testing.T) {
	invalidValue := "invalid-value"
	os.Setenv("LOG_TRACE_CALLER", invalidValue)

	config, err := lgr.NewConfig()

	assert.Equal(t, lgr.Config{}, config, "Should logger config be empty")
	assert.Error(t, err, "Should be error creating logger config")
	assert.Contains(
		t,
		err.Error(),
		fmt.Sprintf(
			`parse error on field "TraceCaller" of type "bool": strconv.ParseBool: parsing "%s": invalid syntax`,
			invalidValue),
		"Logger trace caller value should be bool type",
	)
}
