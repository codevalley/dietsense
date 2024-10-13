package tests

import (
	"dietsense/internal/api"
	"dietsense/internal/services"
	"dietsense/pkg/config"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzeEndpointWithMock(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Create a mock configuration
	mockConfig := &config.AppConfig{
		ServiceType:            "mock",
		MockServiceType:        "default",
		DefaultAnalyzerService: "mock",
	}

	factory := services.NewServiceFactory(mockConfig)
	api.SetupRoutes(router, factory, nil)

	req, _ := http.NewRequest("POST", "/api/v1/analyze", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected HTTP status code 200")

	body, _ := io.ReadAll(resp.Body)
	var result services.AnalysisResult
	err := json.Unmarshal(body, &result)

	assert.NoError(t, err, "Should be able to unmarshal the response")
	assert.NotNil(t, result, "The result should not be nil")
	assert.Equal(t, "mock", result.Service, "Service should be 'mock'")
	assert.NotEmpty(t, result.Summary, "Summary should not be empty")
	assert.NotNil(t, result.NutritionInfo, "NutritionInfo should not be nil")
	assert.Greater(t, result.Confidence, 0.0, "Confidence should be greater than 0")
}
