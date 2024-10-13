package services

import (
	"context"
	"dietsense/pkg/config"
	"dietsense/pkg/logging"
	"dietsense/pkg/utils"
	"encoding/json"
	"fmt"
	"io"

	"github.com/liushuangls/go-anthropic/v2"
)

type ClaudeService struct {
	APIKey    string
	ModelType string
	Config    *config.AppConfig
}

func NewClaudeService(apiKey string, modelType string, config *config.AppConfig) *ClaudeService {
	return &ClaudeService{
		APIKey:    apiKey,
		ModelType: modelType,
		Config:    config,
	}
}

func (s *ClaudeService) ClassifyImage(file io.Reader) (InputType, error) {
	client := anthropic.NewClient(s.APIKey)
	logging.Log.Info("Claude Service: Classifying image, model: " + s.ModelType)

	imageData, err := io.ReadAll(file)
	if err != nil {
		return InputTypeUnknown, fmt.Errorf("failed to read image file: %w", err)
	}

	prompt := s.Config.GetPrompt("claude", "classify_image_prompt")

	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: s.ModelType,
		Messages: []anthropic.Message{
			{
				Role: anthropic.RoleUser,
				Content: []anthropic.MessageContent{
					anthropic.NewImageMessageContent(anthropic.MessageContentImageSource{
						Type:      "base64",
						MediaType: "image/jpeg",
						Data:      imageData,
					}),
					anthropic.NewTextMessageContent(prompt),
				},
			},
		},
		MaxTokens: 100,
	})
	if err != nil {
		return InputTypeUnknown, fmt.Errorf("classification error: %w", err)
	}

	content := resp.Content[0].Text
	switch *content {
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

func (s *ClaudeService) AnalyzeFood(file io.Reader, userContext string, inputType InputType) (*AnalysisResult, error) {
	client := anthropic.NewClient(s.APIKey)
	logging.Log.Info("Claude Service: Analyzing food, model: " + s.ModelType)

	imageData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	var promptString string
	switch inputType {
	case InputTypeFoodImage:
		promptString = "food_image_prompt"
	case InputTypeNutritionLabel:
		promptString = "nutrition_label_prompt"
	case InputTypeBarcode:
		promptString = "barcode_prompt"
	default:
		promptString = "default_image_prompt"
	}
	prompt := s.Config.GetPrompt("claude", promptString)
	jsonFormatInstruction := s.Config.GetPrompt("claude", "json_format_instruction")
	fullContext := fmt.Sprintf("%s\n%s\n%s", prompt, userContext, jsonFormatInstruction)

	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: s.ModelType,
		Messages: []anthropic.Message{
			{
				Role: anthropic.RoleUser,
				Content: []anthropic.MessageContent{
					anthropic.NewImageMessageContent(anthropic.MessageContentImageSource{
						Type:      "base64",
						MediaType: "image/jpeg",
						Data:      imageData,
					}),
					anthropic.NewTextMessageContent(fullContext),
				},
			},
		},
		MaxTokens: 1000,
	})
	if err != nil {
		return nil, fmt.Errorf("analysis error: %w", err)
	}

	return s.parseClaudeResponse(&resp, inputType)
}

func (s *ClaudeService) AnalyzeFoodText(userContext string) (*AnalysisResult, error) {
	client := anthropic.NewClient(s.APIKey)
	logging.Log.Info("Claude Service: Analyzing food description, model: " + s.ModelType)

	fullContext := fmt.Sprintf("%s\n%s", userContext, "Analyze this food description and provide nutritional information in JSON format with 'summary' and 'nutrition' fields.")

	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: s.ModelType,
		Messages: []anthropic.Message{
			{
				Role: anthropic.RoleUser,
				Content: []anthropic.MessageContent{
					anthropic.NewTextMessageContent(fullContext),
				},
			},
		},
		MaxTokens: 1000,
	})
	if err != nil {
		return nil, fmt.Errorf("text analysis error: %w", err)
	}

	return s.parseClaudeResponse(&resp, InputTypeText)
}

func (s *ClaudeService) parseClaudeResponse(resp *anthropic.MessagesResponse, inputType InputType) (*AnalysisResult, error) {
	content := resp.Content[0].Text
	logging.Log.Infof("Claude Response: %s", *content)
	normalizedContent := utils.NormalizeJSON(*content)

	if normalizedContent == "" {
		return nil, fmt.Errorf("failed to extract valid JSON content")
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(normalizedContent), &data); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %w", err)
	}

	result := &AnalysisResult{
		NutritionInfo: make(map[string]interface{}),
		Service:       "claude",
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
