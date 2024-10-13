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

// ClassifyImage implements the ImageClassifier interface.
func (s *MockImageAnalysisService) ClassifyImage(file io.Reader) (InputType, error) {
	logging.Log.Info("Mock Service: Classifying image, model: " + s.ModelType)
	return InputTypeFoodImage, nil // Always return FoodImage for simplicity
}

// AnalyzeFood implements the FoodAnalysisService interface.
func (s *MockImageAnalysisService) AnalyzeFood(file io.Reader, context string, inputType InputType) (*AnalysisResult, error) {
	logging.Log.Info("Mock Service: Analyzing food image, model: " + s.ModelType)
	return &AnalysisResult{
		NutritionInfo: mockNutritionData[0], // Just use the first item for simplicity
		Summary:       "This is a mock summary for testing purposes. It describes a healthy seaweed salad containing wakame, sprouts, sesame seeds, and grated carrots or daikon radish.",
		Confidence:    0.8,
		InputType:     inputType,
		Service:       "mock",
	}, nil
}

// AnalyzeFoodText implements the FoodAnalysisService interface.
func (s *MockImageAnalysisService) AnalyzeFoodText(context string) (*AnalysisResult, error) {
	logging.Log.Info("Mock Service: Analyzing food description, model: " + s.ModelType)
	return &AnalysisResult{
		NutritionInfo: mockNutritionData[0], // Just use the first item for simplicity
		Summary:       "This is a mock summary for text-only analysis. It describes a hypothetical meal based on the provided context.",
		Confidence:    0.8,
		InputType:     InputTypeText,
		Service:       "mock",
	}, nil
}
