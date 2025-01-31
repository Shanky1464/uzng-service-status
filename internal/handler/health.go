package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthLiveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetHealthReadiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
