package services

import (
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

type NutritionDetail struct {
	Component  string  `json:"component"`
	Value      string  `json:"value"`
	Unit       string  `json:"unit"`
	Confidence float64 `json:"confidence"`
}

type Dietsense struct {
	Summary   string            `json:"summary"`
	Nutrition []NutritionDetail `json:"nutrition"`
}

func NewOpenAIService(apiKey string, modelType string) *OpenAIService {
	return &OpenAIService{
		APIKey:    apiKey,
		ModelType: modelType,
	}
}

func (s *OpenAIService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	encodedImage := utils.EncodeToBase64(file)
	logging.Log.Info("OpenAI Service: Analyzing food image, model: " + s.getModel())
	payload := s.createPayload(encodedImage, context)
	responseData, err := utils.SendHTTPRequest("https://api.openai.com/v1/chat/completions", s.APIKey, payload)
	if err != nil {
		return nil, err
	}

	return parseOpenAIResponse(responseData)
}

func (s *OpenAIService) getModel() string {
	switch s.ModelType {
	case "fast":
		return "gpt-4-vision-preview" //"gpt-3.5-turbo"
	case "accurate":
		return "gpt-4-vision-preview" //"gpt-4"
	default:
		return "gpt-4-vision-preview" //"gpt-4-turbo-preview"
	}
}

func (s *OpenAIService) createPayload(encodedImage, context string) map[string]interface{} {
	return map[string]interface{}{
		"model": s.getModel(),
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

func parseOpenAIResponse(response map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	content := response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	normalizedContent := utils.NormalizeJSON(content)

	if normalizedContent == "" {
		return nil, fmt.Errorf("failed to extract valid JSON content")
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(normalizedContent), &data); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %w", err)
	}

	if dietsense, ok := data["dietsense"].([]interface{}); ok && len(dietsense) > 0 {
		if summary, ok := dietsense[0].(map[string]interface{})["summary"].(string); ok {
			result["summary"] = summary
		}
	}

	if nutrition, ok := data["nutrition"].([]interface{}); ok {
		nutritionDetails := make([]map[string]interface{}, len(nutrition))
		for i, item := range nutrition {
			if detail, ok := item.(map[string]interface{}); ok {
				nutritionDetails[i] = map[string]interface{}{
					"component":  detail["component"],
					"value":      detail["value"],
					"unit":       detail["unit"],
					"confidence": detail["confidence"],
				}
			}
		}
		result["nutrition"] = nutritionDetails
	}

	if id, ok := response["id"].(string); ok {
		result["id"] = id
	}

	if usage, ok := response["usage"].(map[string]interface{}); ok {
		if promptTokens, ok := usage["prompt_tokens"].(float64); ok {
			result["prompt tokens"] = int(promptTokens)
		}
	}
	result["service"] = "openAI"
	return result, nil
}

func (s *OpenAIService) AnalyzeFoodText(context string) (map[string]interface{}, error) {
	payload := s.createTextPayload(context)
	logging.Log.Info("OpenAI Service: Analyzing food description, model: " + s.getModel())
	responseData, err := utils.SendHTTPRequest("https://api.openai.com/v1/chat/completions", s.APIKey, payload)
	if err != nil {
		return nil, err
	}

	return parseOpenAIResponse(responseData)
}

func (s *OpenAIService) createTextPayload(context string) map[string]interface{} {
	return map[string]interface{}{
		"model": s.getModel(),
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": context,
			},
		},
		"max_tokens": 4096,
	}
}
