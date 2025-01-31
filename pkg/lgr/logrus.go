package lgr

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var HostnameFunc = os.Hostname

func GetHostname() (string, error) {
	return HostnameFunc()
}

type logFormat struct {
	hostname  string
	formatter logrus.Formatter
}

func (f logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["hostname"] = f.hostname
	return f.formatter.Format(entry)
}

func Setup(config Config) {
	hostname, err := GetHostname()
	if err == nil {
		logrus.SetFormatter(&logFormat{
			hostname:  hostname,
			formatter: &logrus.JSONFormatter{},
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetReportCaller(config.TraceCaller)
	logrus.SetOutput(os.Stderr)

	str, _ := json.Marshal(config)

	switch strings.ToUpper(config.Level) {
	case "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARNING":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.Debugln("Logging configuration", string(str))
}

func New() *logrus.Logger {
	return logrus.New()
}

func StandardLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func Traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
