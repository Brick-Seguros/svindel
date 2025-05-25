package entitystrat

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Passed                 bool     `json:"passed"`
	ExpectedDocumentType   string   `json:"expected_document_type"`
	ExpectedDocumentValue  string   `json:"expected_document_value"`
	Tags                   []string `json:"tags"`
	Comments               string   `json:"comments"`
	ExtractedDocumentType  string   `json:"extracted_document_type"`
	ExtractedDocumentValue string   `json:"extracted_document_value"`
}

func parseResponse(raw string) (response, error) {
	var parsed response
	err := json.Unmarshal([]byte(raw), &parsed)
	if err != nil {
		return response{}, fmt.Errorf("failed to parse entity judge output: %w", err)
	}
	return parsed, nil
}
