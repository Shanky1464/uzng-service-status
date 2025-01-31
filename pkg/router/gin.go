package router

import (
	"uniphore.com/platform-hello-world-go/pkg/lgr"

	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func New(config Config) *gin.Engine {
	switch config.Mode {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	accessLogs := lgr.New()
	accessLogs.SetFormatter(lgr.StandardLogger().Formatter)
	accessLogs.SetOutput(lgr.StandardLogger().Out)
	accessLogs.SetReportCaller(false)

	r := gin.New()
	r.Use(ginlogrus.Logger(accessLogs))
	r.Use(gintrace.Middleware(config.APM.Service))
	r.Use(gin.Recovery())

	r.SetTrustedProxies(nil)

	return r
}
