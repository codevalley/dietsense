package tests

import (
	"dietsense/internal/api"
	"dietsense/internal/services"
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
	api.SetupRoutes(router, services.NewMockImageAnalysisService())

	req, _ := http.NewRequest("POST", "/api/v1/analyze", nil) // Modify the request as needed for your endpoint
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Using testify's assert to simplify test validations
	assert.Equal(t, http.StatusOK, resp.Code, "Expected HTTP status code 200")

	// Read the response body
	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	// Assertions can continue using the parsed data
	assert.NotNil(t, data, "The data should not be nil")
	// You can add more detailed checks on the contents of `data` here
}
