package services

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	// Convert the image file to a byte array
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}
	result := make(map[string]interface{})
	// Encode the image in base64
	encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Construct the JSON payload
	payload := map[string]interface{}{
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
						"type":      "image_url",
						"image_url": map[string]string{"url": fmt.Sprintf("data:image/jpeg;base64,%s", encodedImage)},
					},
				},
			},
		},
		"max_tokens": 4096,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request payload: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	// Send the request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP error status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	// Decode the JSON response
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	// Extract content from the "message" field
	content := response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	// Normalize and parse the embedded JSON
	normalizedContent := strings.TrimPrefix(strings.TrimSuffix(content, "\n```"), "```json\n")
	var dietsense map[string][]map[string]interface{}
	if err := json.Unmarshal([]byte(normalizedContent), &dietsense); err != nil {
		return nil, fmt.Errorf("failed to parse embedded JSON: %w", err)
	}

	// Populate result map with formatted data
	summary := dietsense["Dietsense"][0]["Summary"].(string)
	result["summary"] = summary
	nutritionDetails := make(map[string]string)
	for _, item := range dietsense["Dietsense"][0]["Nutrition"].([]interface{}) {
		detail := item.(map[string]interface{})
		key := detail["Component"].(string)
		value := fmt.Sprintf("%v (%v | %v)", detail["Value"], detail["Unit"], detail["Confidence"])
		nutritionDetails[strings.ToLower(key)] = value
	}
	result["nutrition"] = nutritionDetails

	// Additional metadata
	result["id"] = response["id"].(string)
	result["prompt tokens"] = response["usage"].(map[string]interface{})["prompt_tokens"]

	return result, nil
}
