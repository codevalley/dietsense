package services

import "io"

// FoodAnalysisService defines the interface for an image analysis service.
type FoodAnalysisService interface {
	AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error)
}
