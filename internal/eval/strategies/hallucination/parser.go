package hallucinationstrat

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Passed            bool     `json:"passed"`
	HallucinationRisk string   `json:"hallucination_risk"` // low | medium | high
	Tags              []string `json:"tags"`
	Comments          string   `json:"comments"`
}

func parseResponse(raw string) (response, error) {
	var parsed response
	err := json.Unmarshal([]byte(raw), &parsed)
	if err != nil {
		return response{}, fmt.Errorf("failed to parse hallucination judge output: %w", err)
	}
	return parsed, nil
}
