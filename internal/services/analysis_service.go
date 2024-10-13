package services

import (
	"io"
)

// InputType represents the type of input for analysis
type InputType int

const (
	InputTypeUnknown InputType = iota
	InputTypeFoodImage
	InputTypeNutritionLabel
	InputTypeBarcode
	InputTypeText
)

// AnalysisResult represents the standardized result of an analysis
type AnalysisResult struct {
	NutritionInfo map[string]interface{} `json:"nutrition_info"`
	Summary       string                 `json:"summary"`
	Confidence    float64                `json:"confidence"`
	InputType     InputType              `json:"input_type"`
	Service       string                 `json:"service"`
}

// ImageClassifier defines the interface for classifying images
type ImageClassifier interface {
	ClassifyImage(file io.Reader) (InputType, error)
}

// FoodAnalysisService defines the interface for an image analysis service.
type FoodAnalysisService interface {
	ImageClassifier
	AnalyzeFood(file io.Reader, context string, inputType InputType) (*AnalysisResult, error)
	AnalyzeFoodText(context string) (*AnalysisResult, error)
}
