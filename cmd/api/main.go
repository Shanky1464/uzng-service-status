package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"uniphore.com/uzng-service-status/internal/api/app"
	"uniphore.com/uzng-service-status/internal/handler"
	"uniphore.com/uzng-service-status/internal/handler/v1api"
	"uniphore.com/uzng-service-status/pkg/apm"
	"uniphore.com/uzng-service-status/pkg/lgr"
	"uniphore.com/uzng-service-status/pkg/metrics"
	"uniphore.com/uzng-service-status/pkg/router"
)

const (
	V1Path        = "/v1"
	HealthPath    = "/health"
	CheckAPIPath  = "/check-api"
	AuthToken     = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkxYejBQYWVFRERuS3dfcjdkU2s2UyJ9.eyJodHRwczovL2FwaS51bmlwaG9yZS5jb20vdGVuYW50IjoiZmU1NzIxNTYtMjIxMi00NThlLTg3M2EtNmFjNjBiZWJmNjQyIiwiaHR0cHM6Ly9hcGkudW5pcGhvcmUuY29tL29yZ19pZCI6Im9yZ19TcjJkRHJXSTlvaFY5OFJxIiwiaXNzIjoiaHR0cHM6Ly9kZXYtZzJqNW11MnoudXMuYXV0aDAuY29tLyIsInN1YiI6ImtNbXAzcXNKTWN2NjNOZ0k1RXJyS0hkdlJCM0YxNjdmQGNsaWVudHMiLCJhdWQiOiJhcGkudW5pcGhvcmUuY29tIiwiaWF0IjoxNzM4NzQwODI4LCJleHAiOjE3Mzg4MjcyMjgsInNjb3BlIjoicmVhZDp1c2VycyByZWFkOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSB3cml0ZTp1em5nLXByb2Nlc3NpbmctbWV0YWRhdGEiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMiLCJhenAiOiJrTW1wM3FzSk1jdjYzTmdJNUVycktIZHZSQjNGMTY3ZiIsInBlcm1pc3Npb25zIjpbInJlYWQ6dXNlcnMiLCJyZWFkOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSIsIndyaXRlOnV6bmctcHJvY2Vzc2luZy1tZXRhZGF0YSJdfQ.VLQbFO8z2DqaA6DEXijVM6piMBzwzXTBzlmofEqJYuYpfoiQb9sx7eQIuzrQxrlbQhBmvKyLg8lmkklXWI_A7rBSsoTWAbikGIPOFomXeXdDmu4UEjK7rUjDLo2vvNV8fI2o5dlo-oPIjMsDeY6dmBAqn83xFmqVbqQ_HY_ekX7EMSznDu-Jci5r6RMa4rXs0gqPvLkGKEre0vpWAVo4zEbv27tTiWgwu47kvSOUtYPw6i3a8Cg6r0t_WBk_vBYR3b_1iqXjS0NnowcjWIc27O5yHKE-_nJYwfGunCJzmfDFOT7DgpWUfD4ltrtSheX4z7cPen09pKRhG5Ffb9IETA" // Replace with actual token
)

// Response struct for JSON response
type Response struct {
	StatusCodes map[string]int `json:"status_codes"`
	Message     string         `json:"message"`
}

// Calls an external API and returns the HTTP status code
func callExternalAPI(apiURL string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+AuthToken)

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// Gin-compatible handler for checking multiple APIs
func checkAPIHandler(c *gin.Context) {
	apiEndpoints := map[string]string{
		"Gateway_Startup":      "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/startup",
		"Gateway_Readiness":    "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/readiness",
		"Gateway_Liveness":     "https://api.us.cloud.uniphorestaging.com/uzng-gateway/health/liveness",
		"FlowManager_Startup":  "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/startup",
		"FlowManager_Readiness": "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/readiness",
		"FlowManager_Liveness": "https://api.us.cloud.uniphorestaging.com/uzng-flow-manager/health/liveness",
	}

	statusCodes := make(map[string]int)
	for key, url := range apiEndpoints {
		statusCode, _ := callExternalAPI(url)
		statusCodes[key] = statusCode
	}

	response := Response{
		StatusCodes: statusCodes,
		Message:     "API call completed successfully",
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	// Initialize app config
	appConfig, err := app.NewConfig()
	if err != nil {
		lgr.Fatalf("Failed to initialize configuration: %v", err)
	}

	apm.Start()
	defer apm.Stop()

	lgr.Setup(appConfig.Logger)

	metrics, err := metrics.New(appConfig.Metrics)
	if err != nil {
		lgr.Fatalf("Failed to create DataDog metrics client: %s", err)
	}
	defer metrics.Close()

	// Setup Gin router
	router := router.New(appConfig.Router)
	helloWorldHandlerV1 := v1api.NewHelloWorld(metrics)

	// Define routes
	v1 := router.Group(V1Path)
	{
		v1.GET("/hello", helloWorldHandlerV1.Get)
	}

	// Internal health routes
	router.GET(HealthPath+"/liveness", handler.GetHealthLiveness)
	router.GET(HealthPath+"/readiness", handler.GetHealthReadiness)

	// API status checker route
	router.GET(CheckAPIPath, checkAPIHandler)

	// Run server
	log.Printf("Server running on port %d...", appConfig.Router.Port)
	router.Run(fmt.Sprintf(":%d", appConfig.Router.Port))
}
