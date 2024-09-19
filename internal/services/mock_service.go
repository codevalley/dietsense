package services

import (
	"io"
)

// MockImageAnalysisService is a mock implementation of the ImageAnalysisService
// for testing purposes.
type MockImageAnalysisService struct {
	ModelType string
}

// NewMockImageAnalysisService creates a new instance of MockImageAnalysisService.
func NewMockImageAnalysisService(modelType string) *MockImageAnalysisService {
	return &MockImageAnalysisService{
		ModelType: modelType,
	}
}

// AnalyzeImage implements the ImageAnalysisService interface.
// It ignores the input and returns a fixed, predefined result.
func (s *MockImageAnalysisService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	// Return a mock response that matches the sample-response.json structure
	mockResponse := map[string]interface{}{
		"nutrition": []map[string]interface{}{
			{
				"component":  "Calories",
				"confidence": 0.7,
				"unit":       "kcal",
				"value":      70,
			},
			{
				"component":  "Total Fat",
				"confidence": 0.6,
				"unit":       "g",
				"value":      2,
			},
			{
				"component":  "Saturated Fat",
				"confidence": 0.8,
				"unit":       "g",
				"value":      0,
			},
			{
				"component":  "Cholesterol",
				"confidence": 0.9,
				"unit":       "mg",
				"value":      0,
			},
			{
				"component":  "Sodium",
				"confidence": 0.6,
				"unit":       "mg",
				"value":      150,
			},
			{
				"component":  "Total Carbohydrates",
				"confidence": 0.7,
				"unit":       "g",
				"value":      8,
			},
			{
				"component":  "Dietary Fiber",
				"confidence": 0.8,
				"unit":       "g",
				"value":      6,
			},
			{
				"component":  "Sugars",
				"confidence": 0.6,
				"unit":       "g",
				"value":      2,
			},
			{
				"component":  "Protein",
				"confidence": 0.7,
				"unit":       "g",
				"value":      3,
			},
		},
		"service":    "mock",
		"summary":    "This is a mock summary for testing purposes. It describes a healthy seaweed salad containing wakame, sprouts, sesame seeds, and grated carrots or daikon radish.",
		"model_type": s.ModelType,
	}
	return mockResponse, nil
}

func (s *MockImageAnalysisService) AnalyzeFoodText(context string) (map[string]interface{}, error) {
	// Return the same mock response as AnalyzeFood, but with a different summary
	mockResponse := map[string]interface{}{
		"nutrition": []map[string]interface{}{
			// ... (same nutrition data as before)
		},
		"service":    "mock",
		"summary":    "This is a mock summary for text-only analysis. It describes a hypothetical meal based on the provided context.",
		"model_type": s.ModelType,
	}
	return mockResponse, nil
}
