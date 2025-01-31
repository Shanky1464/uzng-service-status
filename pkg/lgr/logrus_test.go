//go:build unit

package lgr_test

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/lgr"
)

func TestNewLogger(t *testing.T) {
	logger := lgr.New()

	assert.NotNil(t, logger, "Should not be nil")
	assert.Equal(t, logrus.InfoLevel, logger.Level, "Should default log level be INFO")
}

func TestSetupLogger(t *testing.T) {
	cases := []lgr.Config{
		{"TRACE", true},
		{"TRACE", false},
		{"DEBUG", true},
		{"DEBUG", false},
		{"INFO", true},
		{"INFO", false},
		{"WARNING", true},
		{"WARNING", false},
		{"ERROR", true},
		{"ERROR", false},
		{"FATAL", true},
		{"FATAL", false},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			lgr.Setup(tc)

			assert.Equal(
				t,
				strings.ToLower(tc.Level),
				lgr.StandardLogger().GetLevel().String(),
				"Should match log level",
			)
			assert.Equal(
				t,
				tc.TraceCaller,
				lgr.StandardLogger().ReportCaller,
				"Should match report caller setting",
			)
		})
	}

}

func TestSetupLoggerDefaultConfig(t *testing.T) {
	config := lgr.Config{}

	lgr.Setup(config)

	assert.Equal(t, logrus.InfoLevel, lgr.StandardLogger().GetLevel(), "Should log level set to INFO by default")
	assert.False(t, lgr.StandardLogger().ReportCaller, "Should enable report caller by default")
}

func TestTraceLogging(t *testing.T) {
	config := lgr.Config{Level: "TRACE"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Trace("Trace", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.TraceLevel, entry.Level, "Should log level set to TRACE")
	assert.Equal(t, "Tracemessage", entry.Message, "Should match log message")
}

func TestTracelnLogging(t *testing.T) {
	config := lgr.Config{Level: "TRACE"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Traceln("Trace", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.TraceLevel, entry.Level, "Should log level set to TRACE")
	assert.Equal(t, "Trace message", entry.Message, "Should match log message")
}

func TestTracefLogging(t *testing.T) {
	config := lgr.Config{Level: "TRACE"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Tracef("Trace message with %s", "formatting")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.TraceLevel, entry.Level, "Should log level set to TRACE")
	assert.Equal(t, "Trace message with formatting", entry.Message, "Should match log message")
}

func TestDebugLogging(t *testing.T) {
	config := lgr.Config{Level: "DEBUG"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Debug("Debug message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.DebugLevel, entry.Level, "Should log level set to DEBUG")
	assert.Equal(t, "Debug message", entry.Message, "Should match log message")
}

func TestDebuglnLogging(t *testing.T) {
	config := lgr.Config{Level: "DEBUG"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Debugln("Debug", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.DebugLevel, entry.Level, "Should log level set to DEBUG")
	assert.Equal(t, "Debug message", entry.Message, "Should match log message")
}

func TestDebugfLogging(t *testing.T) {
	config := lgr.Config{Level: "DEBUG"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Debugf("Debug message with %s", "formatting")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.DebugLevel, entry.Level, "Should log level set to DEBUG")
	assert.Equal(t, "Debug message with formatting", entry.Message, "Should match log message")
}

func TestInfoLogging(t *testing.T) {
	config := lgr.Config{Level: "INFO"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Info("Info", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.InfoLevel, entry.Level, "Should log level set to INFO")
	assert.Equal(t, "Infomessage", entry.Message, "Should match log message")
}

func TestInfolnLogging(t *testing.T) {
	config := lgr.Config{Level: "INFO"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Infoln("Info", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.InfoLevel, entry.Level, "Should log level set to INFO")
	assert.Equal(t, "Info message", entry.Message, "Should match log message")
}

func TestInfofLogging(t *testing.T) {
	config := lgr.Config{Level: "INFO"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Infof("Info message with %s", "formatting")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.InfoLevel, entry.Level, "Should log level set to INFO")
	assert.Equal(t, "Info message with formatting", entry.Message, "Should match log message")
}

func TestWarningLogging(t *testing.T) {
	config := lgr.Config{Level: "WARNING"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Warn("Warning", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.WarnLevel, entry.Level, "Should log level set to WARNING")
	assert.Equal(t, "Warningmessage", entry.Message, "Should match log message")
}

func TestWarninglnLogging(t *testing.T) {
	config := lgr.Config{Level: "WARNING"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Warnln("Warning", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.WarnLevel, entry.Level, "Should log level set to WARNING")
	assert.Equal(t, "Warning message", entry.Message, "Should match log message")
}

func TestWarningfLogging(t *testing.T) {
	config := lgr.Config{Level: "WARNING"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Warnf("Warning message with %s", "formatting")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.WarnLevel, entry.Level, "Should log level set to WARNING")
	assert.Equal(t, "Warning message with formatting", entry.Message, "Should match log message")
}

func TestErrorLogging(t *testing.T) {
	config := lgr.Config{Level: "ERROR"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Error("Error", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.ErrorLevel, entry.Level, "Should log level set to ERROR")
	assert.Equal(t, "Errormessage", entry.Message, "Should match log message")
}

func TestErrorlnLogging(t *testing.T) {
	config := lgr.Config{Level: "ERROR"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Errorln("Error", "message")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.ErrorLevel, entry.Level, "Should log level set to ERROR")
	assert.Equal(t, "Error message", entry.Message, "Should match log message")
}

func TestErrorfLogging(t *testing.T) {
	config := lgr.Config{Level: "ERROR"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	lgr.Errorf("Error message with %s", "formatting")
	entry := hook.LastEntry()

	assert.Equal(t, logrus.ErrorLevel, entry.Level, "Should log level set to ERROR")
	assert.Equal(t, "Error message with formatting", entry.Message, "Should match log message")
}

func TestFatalLogging(t *testing.T) {
	config := lgr.Config{Level: "FATAL"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	fatal := false
	lgr.StandardLogger().ExitFunc = func(int) { fatal = true }
	defer func() { lgr.StandardLogger().ExitFunc = nil }()

	lgr.Fatal("Fatal", "message")
	entry := hook.LastEntry()

	assert.True(t, fatal, "Should update the fatal variable to true")
	assert.Equal(t, logrus.FatalLevel, entry.Level, "Should log level set to FATAL")
	assert.Equal(t, "Fatalmessage", entry.Message, "Should match log message")
}

func TestFatallnLogging(t *testing.T) {
	config := lgr.Config{Level: "FATAL"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	fatal := false
	lgr.StandardLogger().ExitFunc = func(int) { fatal = true }
	defer func() { lgr.StandardLogger().ExitFunc = nil }()

	lgr.Fatalln("Fatal", "message")
	entry := hook.LastEntry()

	assert.True(t, fatal, "Should update the fatal variable to true")
	assert.Equal(t, logrus.FatalLevel, entry.Level, "Should log level set to FATAL")
	assert.Equal(t, "Fatal message", entry.Message, "Should match log message")
}

func TestFatalfLogging(t *testing.T) {
	config := lgr.Config{Level: "FATAL"}
	lgr.Setup(config)
	hook := test.NewLocal(lgr.StandardLogger())

	fatal := false
	lgr.StandardLogger().ExitFunc = func(int) { fatal = true }
	defer func() { lgr.StandardLogger().ExitFunc = nil }()

	lgr.Fatalf("Fatal message with %s", "formatting")
	entry := hook.LastEntry()

	assert.True(t, fatal, "Should update the fatal variable to true")
	assert.Equal(t, logrus.FatalLevel, entry.Level, "Should log level set to FATAL")
	assert.Equal(t, "Fatal message with formatting", entry.Message, "Should match log message")
}

func TestHostname(t *testing.T) {
	lgr.HostnameFunc = func() (string, error) {
		return "some-hostname", nil
	}
	defer func() {
		lgr.HostnameFunc = os.Hostname
	}()

	config := lgr.Config{Level: "INFO"}
	lgr.Setup(config)
	var buf bytes.Buffer
	lgr.StandardLogger().SetOutput(&buf)

	lgr.Info("some message")

	assert.Contains(t, buf.String(), `"hostname":"some-hostname"`)
}

func TestHostnameError(t *testing.T) {
	lgr.HostnameFunc = func() (string, error) {
		return "", errors.New("Error")
	}
	defer func() {
		lgr.HostnameFunc = os.Hostname
	}()

	config := lgr.Config{Level: "INFO"}
	lgr.Setup(config)
	var buf bytes.Buffer
	lgr.StandardLogger().SetOutput(&buf)

	lgr.Info("some message")

	assert.NotContains(t, buf.String(), `"hostname":`)
}
