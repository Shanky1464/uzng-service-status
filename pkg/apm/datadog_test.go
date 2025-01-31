//go:build unit

package apm_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	"uniphore.com/platform-hello-world-go/internal/handler"
	"uniphore.com/platform-hello-world-go/pkg/router"
)

func TestAPM(t *testing.T) {
	mt := mocktracer.Start()
	defer mt.Stop()

	r := router.New(router.Config{Mode: "test"})
	r.GET("/health/liveness", handler.GetHealthLiveness)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/liveness", nil)
	r.ServeHTTP(w, req)

	spans := mt.FinishedSpans()
	assert.Len(t, spans, 1, "Should number of spans be equal to 1")

	assert.Equal(t, "http.request", spans[0].OperationName(), "Should match operation name")
	assert.Equal(t, "200", spans[0].Tags()["http.status_code"], "Should match status code")
	assert.Equal(t, "GET", spans[0].Tags()["http.method"], "Should match HTTP method")
	assert.Equal(t, "/health/liveness", spans[0].Tags()["http.url"], "Should match HTTP URL")
	assert.Equal(t, "server", spans[0].Tags()["span.kind"], "Should match span kind")
	assert.Equal(t, "web", spans[0].Tags()["span.type"], "Should match span type")
}
