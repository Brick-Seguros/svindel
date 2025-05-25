package relevancestrat

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Passed         bool     `json:"passed"`
	RelevanceScore int      `json:"relevance_score"` // 0-10
	Tags           []string `json:"tags"`
	Comments       string   `json:"comments"`
}

func parseResponse(raw string) (response, error) {
	var parsed response
	err := json.Unmarshal([]byte(raw), &parsed)
	if err != nil {
		return response{}, fmt.Errorf("failed to parse relevance judge output: %w", err)
	}
	return parsed, nil
}
