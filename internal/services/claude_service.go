package services

import (
	"context"
	"dietsense/pkg/logging"
	"dietsense/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/liushuangls/go-anthropic/v2"
)

type ClaudeService struct {
	APIKey string
}

func NewClaudeService(apiKey string) *ClaudeService {
	return &ClaudeService{
		APIKey: apiKey,
	}
}

func (s *ClaudeService) AnalyzeFood(file io.Reader, userContext string) (map[string]interface{}, error) {
	client := anthropic.NewClient(s.APIKey)

	imageData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3Opus20240229,
		Messages: []anthropic.Message{
			{
				Role: anthropic.RoleUser,
				Content: []anthropic.MessageContent{
					anthropic.NewImageMessageContent(anthropic.MessageContentImageSource{
						Type:      "base64",
						MediaType: "image/jpeg",
						Data:      imageData,
					}),
					anthropic.NewTextMessageContent(userContext),
				},
			},
		},
		MaxTokens: 1000,
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			return nil, fmt.Errorf("messages error, type: %s, message: %s", e.Type, e.Message)
		} else {
			return nil, fmt.Errorf("messages error: %w", err)
		}
	}

	return parseClaudeResponse(&resp)
}

func parseClaudeResponse(resp *anthropic.MessagesResponse) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	content := resp.Content[0].Text
	logging.Log.Infof("Claude Response: %s ", *content)
	normalizedContent := utils.NormalizeJSON(*content)

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

	return result, nil
}
