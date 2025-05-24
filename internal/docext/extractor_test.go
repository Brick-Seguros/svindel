package docext

import (
	"testing"

	"svindel/internal/shared"
)

func TestOpenAiExtractor_RegexExtraction(t *testing.T) {
	extractor := &OpenAiExtractor{}

	tests := []struct {
		name     string
		input    string
		expected shared.ExtractionResult
	}{
		{
			name:  "Valid CPF",
			input: "Check CPF 097.191.979-86",
			expected: shared.ExtractionResult{
				Document:     "097.191.979-86",
				DocumentType: shared.DocTypeCPF,
				IsQuestion:   false,
			},
		},
		{
			name:  "Valid CNPJ",
			input: "Check CNPJ 12.345.678/0001-90",
			expected: shared.ExtractionResult{
				Document:     "12.345.678/0001-90",
				DocumentType: shared.DocTypeCNPJ,
				IsQuestion:   false,
			},
		},
		{
			name:  "Valid Plate",
			input: "Check plate ABC1D23",
			expected: shared.ExtractionResult{
				Document:     "ABC1D23",
				DocumentType: shared.DocTypePlate,
				IsQuestion:   false,
			},
		},
		{
			name:  "No document (question fallback)",
			input: "How does this work?",
			expected: shared.ExtractionResult{
				Document:     "",
				DocumentType: shared.DocTypeNone,
				IsQuestion:   true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractor.Extract(tt.input)

			if result.Document != tt.expected.Document {
				t.Errorf("Expected document %q, got %q", tt.expected.Document, result.Document)
			}

			if result.DocumentType != tt.expected.DocumentType {
				t.Errorf("Expected doc type %q, got %q", tt.expected.DocumentType, result.DocumentType)
			}

			if result.IsQuestion != tt.expected.IsQuestion {
				t.Errorf("Expected IsQuestion %v, got %v", tt.expected.IsQuestion, result.IsQuestion)
			}
		})
	}
}
