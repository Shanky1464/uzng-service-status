//go:build unit

package v1api_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/internal/handler/v1api"
	"uniphore.com/platform-hello-world-go/pkg/metrics"
	"uniphore.com/platform-hello-world-go/pkg/router"
)

func TestGetHelloWorldWithNameParam(t *testing.T) {
	metricsStub, _ := metrics.New(metrics.Config{})

	r := router.New(router.Config{Mode: "test"})
	h := v1api.NewHelloWorld(metricsStub)
	r.GET("/", h.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?name=foobar", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should match HTTP status code")
	assert.JSONEq(t, `{"helloworld": "foobar"}`, w.Body.String(), "Should match response body")
}

func TestGetHelloWorldWithNameAndLastNameParams(t *testing.T) {
	metricsStub, _ := metrics.New(metrics.Config{})

	r := router.New(router.Config{Mode: "test"})
	h := v1api.NewHelloWorld(metricsStub)
	r.GET("/", h.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?name=foo&lastname=bar", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Should match HTTP status code")
	assert.JSONEq(t, `{"helloworld": "foo bar"}`, w.Body.String(), "Should match response body")
}

func TestGetHelloWorldWithMissingRequiredQueryParams(t *testing.T) {
	cases := []string{
		"/",
		"/?lastname=foobar",
	}

	metricsStub, _ := metrics.New(metrics.Config{})

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := router.New(router.Config{Mode: "test"})
			h := v1api.NewHelloWorld(metricsStub)
			r.GET("/", h.Get)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 422, w.Code, "Should match HTTP status code")
			assert.JSONEq(
				t,
				`{
					"code": 422,
					"message": "Request validation error"
				}`,
				w.Body.String(),
				"Should match response body",
			)
		})
	}
}

func TestGetHelloWorldWithInvalidQueryParams(t *testing.T) {
	cases := []string{
		"/?name=abc123",
		"/?lastname=abc123",
		"/?name=!foobar",
		"/?name=foo!bar",
		"/?name=foobar!",
		"/?name=foo&lastname=bar123",
		"/?name=foo123&lastname=bar",
		"/?name=foo123&lastname=bar123",
	}

	metricsStub, _ := metrics.New(metrics.Config{})

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := router.New(router.Config{Mode: "test"})
			h := v1api.NewHelloWorld(metricsStub)
			r.GET("/", h.Get)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tc, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, 422, w.Code, "Incorrect HTTP status code")
			assert.JSONEq(t,
				`{
					"code": 422,
					"message": "Request validation error"
				}`,
				w.Body.String(),
				"Should match response body",
			)
		})
	}
}
