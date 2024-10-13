package services

import (
	"dietsense/pkg/config"
	"dietsense/pkg/logging"
	"dietsense/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
)

type OpenAIService struct {
	APIKey    string
	ModelType string
}

func NewOpenAIService(apiKey string, modelType string) *OpenAIService {
	return &OpenAIService{
		APIKey:    apiKey,
		ModelType: modelType,
	}
}

func (s *OpenAIService) ClassifyImage(file io.Reader) (InputType, error) {
	encodedImage := utils.EncodeToBase64(file)
	logging.Log.Info("OpenAI Service: Classifying image, model: " + s.ModelType)

	prompt := config.Config.ClassifyImagePrompt
	payload := s.createPayload(encodedImage, prompt)

	responseData, err := utils.SendHTTPRequest("https://api.openai.com/v1/chat/completions", s.APIKey, payload)
	if err != nil {
		return InputTypeUnknown, fmt.Errorf("failed to classify image: %w", err)
	}

	content := responseData["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	switch content {
	case "food photo":
		return InputTypeFoodImage, nil
	case "nutrition label":
		return InputTypeNutritionLabel, nil
	case "barcode":
		return InputTypeBarcode, nil
	default:
		return InputTypeUnknown, nil
	}
}

func (s *OpenAIService) AnalyzeFood(file io.Reader, context string, inputType InputType) (*AnalysisResult, error) {
	encodedImage := utils.EncodeToBase64(file)
	logging.Log.Info("OpenAI Service: Analyzing food, model: " + s.ModelType)

	var prompt string
	switch inputType {
	case InputTypeFoodImage:
		prompt = config.Config.FoodImagePrompt
	case InputTypeNutritionLabel:
		prompt = config.Config.NutritionLabelPrompt
	case InputTypeBarcode:
		prompt = config.Config.BarcodePrompt
	default:
		prompt = config.Config.DefaultImagePrompt
	}

	fullContext := fmt.Sprintf("%s\n%s\n%s", prompt, context, "Provide the response in JSON format with 'summary' and 'nutrition' fields.")

	payload := s.createPayload(encodedImage, fullContext)
	responseData, err := utils.SendHTTPRequest("https://api.openai.com/v1/chat/completions", s.APIKey, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze food: %w", err)
	}

	return s.parseOpenAIResponse(responseData, inputType)
}

func (s *OpenAIService) AnalyzeFoodText(context string) (*AnalysisResult, error) {
	logging.Log.Info("OpenAI Service: Analyzing food description, model: " + s.ModelType)
	fullContext := fmt.Sprintf("%s\n%s", context, "Analyze this food description and provide nutritional information in JSON format with 'summary' and 'nutrition' fields.")
	payload := s.createTextPayload(fullContext)
	responseData, err := utils.SendHTTPRequest("https://api.openai.com/v1/chat/completions", s.APIKey, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze food text: %w", err)
	}

	return s.parseOpenAIResponse(responseData, InputTypeText)
}

func (s *OpenAIService) parseOpenAIResponse(response map[string]interface{}, inputType InputType) (*AnalysisResult, error) {
	content := response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	normalizedContent := utils.NormalizeJSON(content)

	if normalizedContent == "" {
		return nil, fmt.Errorf("failed to extract valid JSON content")
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(normalizedContent), &data); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %w", err)
	}

	result := &AnalysisResult{
		NutritionInfo: make(map[string]interface{}),
		Service:       "openAI",
		InputType:     inputType,
	}

	if summary, ok := data["summary"].(string); ok {
		result.Summary = summary
	}

	if nutrition, ok := data["nutrition"].([]interface{}); ok {
		for _, item := range nutrition {
			if detail, ok := item.(map[string]interface{}); ok {
				component := detail["component"].(string)
				result.NutritionInfo[component] = map[string]interface{}{
					"value":      detail["value"],
					"unit":       detail["unit"],
					"confidence": detail["confidence"],
				}
			}
		}
	}

	result.Confidence = 0.8 // Default confidence

	return result, nil
}

func (s *OpenAIService) createPayload(encodedImage, context string) map[string]interface{} {
	return map[string]interface{}{
		"model": s.ModelType,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": context,
					},
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/jpeg;base64,%s", encodedImage),
						},
					},
				},
			},
		},
		"max_tokens": 4096,
	}
}

func (s *OpenAIService) createTextPayload(context string) map[string]interface{} {
	return map[string]interface{}{
		"model": s.ModelType,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": context,
			},
		},
		"max_tokens": 4096,
	}
}
