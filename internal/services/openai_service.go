package services

import (
	"bytes"
	"crypto/tls"
	"dietsense/pkg/logging"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type OpenAIService struct {
	APIKey string
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

func NewOpenAIService(apiKey string) *OpenAIService {
	return &OpenAIService{
		APIKey: apiKey,
	}
}

func (s *OpenAIService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	encodedImage := encodeToBase64(file)
	payload := createPayload(encodedImage, context)
	responseData, err := sendHTTPRequest(s.APIKey, payload)
	if err != nil {
		return nil, err
	}

	return parseResponse(responseData)
}

func encodeToBase64(file io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func createPayload(encodedImage, context string) map[string]interface{} {
	return map[string]interface{}{
		"model": "gpt-4-vision-preview",
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

func sendHTTPRequest(apiKey string, payload map[string]interface{}) (map[string]interface{}, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}
	logging.Log.Info("Response: ", response)
	return response, nil
}

func parseResponse(response map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	content := response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	normalizedContent := normalizeJSON(content)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(normalizedContent), &data); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %w", err)
	}

	if summary := extractSummary(data); summary != "" {
		result["summary"] = summary
	}

	if nutritionDetails := extractNutritionDetails(data); len(nutritionDetails) > 0 {
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

	return result, nil
}

func extractSummary(data map[string]interface{}) string {
	if dietsense, ok := data["Dietsense"].([]interface{}); ok {
		for _, item := range dietsense {
			if summary, ok := item.(map[string]interface{})["Summary"].(string); ok {
				return summary
			}
		}
	}
	return ""
}

func extractNutritionDetails(data map[string]interface{}) map[string]string {
	nutritionDetails := make(map[string]string)

	if nutrition, ok := data["Nutrition"].([]interface{}); ok {
		for _, item := range nutrition {
			if detail, ok := item.(map[string]interface{}); ok {
				if component, ok := detail["Component"].(string); ok {
					key := strings.ToLower(component)
					value := fmt.Sprintf("%v (%v | %v)", detail["Value"], detail["Unit"], detail["Confidence"])
					nutritionDetails[key] = value
				}
			}
		}
	} else if dietsense, ok := data["Dietsense"].([]interface{}); ok {
		for _, item := range dietsense {
			if nutrition, ok := item.(map[string]interface{})["Nutrition"].([]interface{}); ok {
				for _, item := range nutrition {
					if detail, ok := item.(map[string]interface{}); ok {
						if component, ok := detail["Component"].(string); ok {
							key := strings.ToLower(component)
							value := fmt.Sprintf("%v (%v | %v)", detail["Value"], detail["Unit"], detail["Confidence"])
							nutritionDetails[key] = value
						}
					}
				}
			}
		}
	}

	return nutritionDetails
}
func normalizeJSON(content string) string {
	// Remove code block markers and trim whitespace
	re := regexp.MustCompile(`(?s)^(?:\x60{3}json\n|\x60{3}\n)?(.*?)(?:\n\x60{3})?$`)
	normalizedContent := strings.TrimSpace(re.ReplaceAllString(content, "$1"))

	// Remove any trailing newline characters after the JSON object
	re = regexp.MustCompile(`(}\s*)\n+$`)
	normalizedContent = re.ReplaceAllString(normalizedContent, "$1")

	return normalizedContent
}
