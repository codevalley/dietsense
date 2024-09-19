package services

import (
	"dietsense/pkg/logging"
	"io"
)

// MockImageAnalysisService is a mock implementation of the ImageAnalysisService
// for testing purposes.
type MockImageAnalysisService struct {
	ModelType string
}

// mockNutritionData is a constant representing the mock nutrition data
var mockNutritionData = []map[string]interface{}{
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
}

// NewMockImageAnalysisService creates a new instance of MockImageAnalysisService.
func NewMockImageAnalysisService(modelType string) *MockImageAnalysisService {
	return &MockImageAnalysisService{
		ModelType: modelType,
	}
}

// getMockResponse returns a mock response with the given summary
func (s *MockImageAnalysisService) getMockResponse(summary string) map[string]interface{} {
	return map[string]interface{}{
		"nutrition":  mockNutritionData,
		"service":    "mock",
		"summary":    summary,
		"model_type": s.ModelType,
	}
}

// AnalyzeFood implements the ImageAnalysisService interface.
// It ignores the input and returns a fixed, predefined result.
func (s *MockImageAnalysisService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	logging.Log.Info("Mock Service: Analyzing food image, model: " + s.ModelType)
	return s.getMockResponse("This is a mock summary for testing purposes. It describes a healthy seaweed salad containing wakame, sprouts, sesame seeds, and grated carrots or daikon radish."), nil
}

func (s *MockImageAnalysisService) AnalyzeFoodText(context string) (map[string]interface{}, error) {
	logging.Log.Info("Mock Service: Analyzing food description, model: " + s.ModelType)
	return s.getMockResponse("This is a mock summary for text-only analysis. It describes a hypothetical meal based on the provided context."), nil
}
