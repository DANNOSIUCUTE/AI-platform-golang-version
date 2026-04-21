package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Tách hàm gọi API phức tạp khỏi Controller
func GetClaudeResponse(apiKey, message string) (string, error) {
	if apiKey == "" {
		return "Đây là câu trả lời Mock do bạn chưa cấu hình CLAUDE_API_KEY: " + message, nil
	}

	url := "https://api.anthropic.com/v1/messages"

	payloadData := map[string]interface{}{
		"model":      "claude-3-sonnet-20240229",
		"max_tokens": 1000,
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	payloadBytes, _ := json.Marshal(payloadData)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	req.Header.Add("x-api-key", apiKey)
	req.Header.Add("anthropic-version", "2023-06-01")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Lỗi API Claude: %s", string(body))
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// Bóc tách JSON
	if contentArr, ok := result["content"].([]interface{}); ok && len(contentArr) > 0 {
		if firstContent, ok := contentArr[0].(map[string]interface{}); ok {
			if textObj, ok := firstContent["text"].(string); ok {
				return textObj, nil
			}
		}
	}

	return "", fmt.Errorf("Invalid structure returned from Claude")
}
