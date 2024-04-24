package utils

import (
	"bytes"
	"dietsense/pkg/logging"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func EncodeToBase64(file io.Reader) string {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func NormalizeJSON(content string) string {
	// Find the start and end of the JSON content
	startIndex := strings.Index(content, "{")
	endIndex := strings.LastIndex(content, "}")

	if startIndex != -1 && endIndex != -1 && endIndex >= startIndex {
		// Extract the JSON content
		jsonContent := content[startIndex : endIndex+1]

		// Unescape escaped characters
		jsonContent = strings.ReplaceAll(jsonContent, "\\\"", "\"")
		jsonContent = strings.ReplaceAll(jsonContent, "\\n", "\n")

		return jsonContent
	}

	return ""
}

func SendHTTPRequest(url, apiKey string, payload map[string]interface{}) (map[string]interface{}, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}

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
