package services

import (
	"dietsense/pkg/config"
	"fmt"
	"log"
)

// ServiceFactory manages creation of food analysis services.
type ServiceFactory struct {
	DefaultService string
}

// NewServiceFactory creates a new instance of a service factory.
func NewServiceFactory(defaultService string) *ServiceFactory {
	return &ServiceFactory{DefaultService: defaultService}
}

// GetService returns the food analysis service based on the requested type.
func (f *ServiceFactory) GetService(serviceType string) (FoodAnalysisService, error) {
	if serviceType == "" {
		serviceType = f.DefaultService
	}

	switch serviceType {
	case "openai":
		return NewOpenAIService(config.Config.OpenaiKey, config.Config.ModelType), nil
	case "claude":
		return NewClaudeService(config.Config.ClaudeKey, config.Config.ModelType), nil
	case "mock":
		return NewMockImageAnalysisService(config.Config.ModelType), nil
	default:
		log.Printf("Unknown service type: %s, falling back to default.", serviceType)
		return nil, fmt.Errorf("unknown service type: %s", serviceType)
	}
}
