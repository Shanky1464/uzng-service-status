//go:build unit

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"uniphore.com/platform-hello-world-go/internal/handler"
	"uniphore.com/platform-hello-world-go/pkg/router"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthLiveness(t *testing.T) {
	r := router.New(router.Config{Mode: "test"})
	r.GET("/health/liveness", handler.GetHealthLiveness)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/liveness", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should match HTTP status code")
	assert.JSONEq(t, `{"status": "ok"}`, w.Body.String(), "Should match response body")
}

func TestGetHealthReadiness(t *testing.T) {
	r := router.New(router.Config{Mode: "test"})
	r.GET("/health/readiness", handler.GetHealthReadiness)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/readiness", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should match HTTP status code")
	assert.JSONEq(t, `{"status": "ok"}`, w.Body.String(), "Should match response body")
}
