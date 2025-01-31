//go:build unit

package router_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/pkg/router"
)

func TestNewGin(t *testing.T) {
	cases := []struct {
		nameMode   string
		routerMode string
	}{
		{
			nameMode:   "prod",
			routerMode: gin.ReleaseMode,
		},
		{
			nameMode:   "test",
			routerMode: gin.TestMode,
		},
		{
			nameMode:   "debug",
			routerMode: gin.DebugMode,
		},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			config := router.Config{Mode: tc.nameMode}
			req, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()

			router := router.New(config)
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.routerMode, gin.Mode(), "Should set a router mode")
		})
	}
}

func TestNewGinDefaultMode(t *testing.T) {
	config := router.Config{}
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router := router.New(config)
	router.ServeHTTP(w, req)

	assert.Equal(t, gin.DebugMode, gin.Mode(), "Should use a debug mode by default")
}

func TestNewGinInvalidMode(t *testing.T) {
	config := router.Config{Mode: "invalid-mode"}
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router := router.New(config)
	router.ServeHTTP(w, req)

	assert.Equal(t, gin.DebugMode, gin.Mode(), "Should use a debug mode if defined is invalid")
}
