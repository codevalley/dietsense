package services

import (
	"dietsense/pkg/config"
	"fmt"
)

type ServiceFactory struct {
	Config *config.AppConfig
}

func NewServiceFactory(config *config.AppConfig) *ServiceFactory {
	return &ServiceFactory{Config: config}
}

func (f *ServiceFactory) GetImageClassifierService() (ImageClassifier, error) {
	switch f.Config.ImageClassifierService {
	case "openai":
		return NewOpenAIService(f.Config.OpenaiKey, f.Config.OpenAIModelForClassification), nil
	case "claude":
		return NewClaudeService(f.Config.ClaudeKey, f.Config.ClaudeModelForClassification), nil
	default:
		return nil, fmt.Errorf("unknown image classifier service: %s", f.Config.ImageClassifierService)
	}
}

func (f *ServiceFactory) GetAnalyzerService(inputType InputType) (FoodAnalysisService, error) {
	var serviceType string
	switch inputType {
	case InputTypeBarcode:
		serviceType = f.Config.BarcodeAnalyzerService
	case InputTypeFoodImage:
		serviceType = f.Config.FoodImageAnalyzerService
	case InputTypeNutritionLabel:
		serviceType = f.Config.NutritionLabelAnalyzerService
	case InputTypeText:
		serviceType = f.Config.TextAnalyzerService
	default:
		serviceType = f.Config.DefaultAnalyzerService
	}

	switch serviceType {
	case "openai":
		return NewOpenAIService(f.Config.OpenaiKey, f.Config.OpenAIModelForAnalysis), nil
	case "claude":
		return NewClaudeService(f.Config.ClaudeKey, f.Config.ClaudeModelForAnalysis), nil
	case "mock":
		return NewMockImageAnalysisService(f.Config.MockServiceType), nil
	default:
		return nil, fmt.Errorf("unknown analyzer service: %s", serviceType)
	}
}
