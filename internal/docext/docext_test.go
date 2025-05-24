package docext

import (
	"testing"

	"svindel/internal/shared"
)

// Mock Extractor
type MockExtractor struct {
	ExpectedExtraction shared.ExtractionResult
}

func (m *MockExtractor) Extract(input string) shared.ExtractionResult {
	return m.ExpectedExtraction
}

// Mock Retriever
type MockRetriever struct {
	ExpectedContext shared.ContextResult
}

func (m *MockRetriever) Retrieve(doc string, docType shared.DocType) shared.ContextResult {
	return m.ExpectedContext
}

func TestDocExt_Process_Question(t *testing.T) {
	extractor := &MockExtractor{
		ExpectedExtraction: shared.ExtractionResult{
			Document:     "",
			DocumentType: shared.DocTypeNone,
			IsQuestion:   true,
		},
	}

	retriever := &MockRetriever{}

	docExt := New(extractor, retriever)

	userInput := "How does this work?"

	augPrompt, extraction, context := docExt.Process(userInput)

	if augPrompt != userInput {
		t.Errorf("Expected prompt to be user input. Got %q", augPrompt)
	}

	if !extraction.IsQuestion {
		t.Errorf("Expected extraction.IsQuestion to be true")
	}

	if len(context.Reports) != 0 || len(context.Resources) != 0 {
		t.Errorf("Expected empty context for a question")
	}
}

func TestDocExt_Process_WithDocument(t *testing.T) {
	extractor := &MockExtractor{
		ExpectedExtraction: shared.ExtractionResult{
			Document:     "09719197986",
			DocumentType: shared.DocTypeCPF,
			IsQuestion:   false,
		},
	}

	retriever := &MockRetriever{
		ExpectedContext: shared.ContextResult{
			Reports: []shared.ReportResult{
				{
					ID:        "report-123",
					Name:      "Fraud Check",
					Document:  "09719197986",
					CreatedAt: "2025-05-20T16:57:01.937Z",
				},
			},
			Resources: []shared.Resource{
				{
					ID:         "resource-1",
					Title:      "CPF Validator",
					HelperText: "Check if the CPF is valid.",
				},
			},
		},
	}

	docExt := New(extractor, retriever)

	userInput := "Check this CPF 09719197986"

	augPrompt, extraction, context := docExt.Process(userInput)

	// Check augmented prompt contains key elements
	if augPrompt == userInput {
		t.Errorf("Expected augmented prompt to be different from user input")
	}

	if !contains(augPrompt, "09719197986") {
		t.Errorf("Augmented prompt should contain the document")
	}

	if !contains(augPrompt, "Fraud Check") {
		t.Errorf("Augmented prompt should include report title")
	}

	if !contains(augPrompt, "CPF Validator") {
		t.Errorf("Augmented prompt should include resource title")
	}

	// Check extraction
	if extraction.IsQuestion {
		t.Errorf("Expected extraction.IsQuestion to be false")
	}

	if extraction.Document != "09719197986" {
		t.Errorf("Expected extraction document to be '09719197986'")
	}

	// Check context
	if len(context.Reports) != 1 {
		t.Errorf("Expected 1 report, got %d", len(context.Reports))
	}

	if len(context.Resources) != 1 {
		t.Errorf("Expected 1 resource, got %d", len(context.Resources))
	}
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && (len(substr) == 0 || (len(str) > 0 && (str == substr || (len(str) > len(substr) && (str[0:len(substr)] == substr || contains(str[1:], substr))))))
}
