package services

import (
	"io"
)

// MockImageAnalysisService is a mock implementation of the ImageAnalysisService
// for testing purposes.
type MockImageAnalysisService struct{}

// NewMockImageAnalysisService creates a new instance of MockImageAnalysisService.
func NewMockImageAnalysisService() *MockImageAnalysisService {
	return &MockImageAnalysisService{}
}

// AnalyzeImage implements the ImageAnalysisService interface.
// It ignores the input and returns a fixed, predefined result.
func (s *MockImageAnalysisService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	// Return a fixed mock response
	mockResponse := map[string]interface{}{
		"calories":    500,
		"protein":     "25g",
		"carbs":       "50g",
		"fats":        "10g",
		"description": "This is a fixed mock response for testing purposes.",
		"service":     "mock",
	}
	return mockResponse, nil
}
