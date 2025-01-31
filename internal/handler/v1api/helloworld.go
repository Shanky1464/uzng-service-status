package v1api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"uniphore.com/platform-hello-world-go/pkg/lgr"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
)

type HelloWorld struct {
	metrics *metrics.Metrics
}

func NewHelloWorld(metrics *metrics.Metrics) *HelloWorld {
	SetupValidator()
	return &HelloWorld{metrics: metrics}
}

func (h *HelloWorld) Get(c *gin.Context) {
	type queryParams struct {
		Name     string `form:"name" binding:"name"`
		LastName string `form:"lastname" binding:"lastname"`
	}
	var params queryParams

	if err := c.ShouldBindQuery(&params); err != nil {
		lgr.Error(err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"code":    http.StatusUnprocessableEntity,
			"message": "Request validation error",
		})
		return
	}
	lgr.Debugf("Query parameters have been validated correctly: %v", params)

	h.metrics.Incr("request.counter", []string{}, 1)

	if len(params.LastName) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"helloworld": params.Name + " " + params.LastName,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"helloworld": params.Name,
		})
	}
}
