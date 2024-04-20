package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// OpenAIService implements the ImageAnalysisService for OpenAI.
type OpenAIService struct {
	APIKey string
}

// NewOpenAIService creates a new instance of OpenAIService with the provided API key.
func NewOpenAIService(apiKey string) *OpenAIService {
	return &OpenAIService{
		APIKey: apiKey,
	}
}

func (s *OpenAIService) AnalyzeFood(file io.Reader, context string) (map[string]interface{}, error) {
	// Prepare the URL for the OpenAI API endpoint
	url := "https://api.openai.com/v1/images/analysis" // Ensure this is the correct API endpoint

	// Convert the image file to a byte array
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	// Create the request payload
	payload := map[string]interface{}{
		"image":   buf.Bytes(), // This may need to be base64 encoded depending on the API requirements
		"context": context,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request payload: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP error status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	// Decode the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	return result, nil
}
