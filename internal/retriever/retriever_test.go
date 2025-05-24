package retriever

import (
	"testing"

	"svindel/internal/shared"
)

// Mock Report Service
type MockReportService struct {
	ExpectedReports []shared.ReportResult
}

func (m *MockReportService) GetReportsForDocument(doc string, docType shared.DocType) ([]shared.ReportResult, error) {
	return m.ExpectedReports, nil
}

func TestRetriever_Retrieve(t *testing.T) {
	mockReports := []shared.ReportResult{
		{
			ID:        "report-123",
			Name:      "Fraud Check",
			Document:  "09719197986",
			CreatedAt: "2025-05-20T16:57:01.937Z",
		},
	}

	mockService := &MockReportService{
		ExpectedReports: mockReports,
	}

	retriever := New(mockService)

	result := retriever.Retrieve("09719197986", shared.DocTypeCPF)

	// 🔥 Check reports
	if len(result.Reports) != 1 {
		t.Fatalf("Expected 1 report, got %d", len(result.Reports))
	}

	if result.Reports[0].ID != "report-123" {
		t.Errorf("Expected report ID 'report-123', got '%s'", result.Reports[0].ID)
	}

	// 🔥 Check resources
	if len(result.Resources) != 1 {
		t.Fatalf("Expected 1 resource, got %d", len(result.Resources))
	}

	expectedResourceID := "resource-validate-cpf"
	if result.Resources[0].ID != expectedResourceID {
		t.Errorf("Expected resource ID '%s', got '%s'", expectedResourceID, result.Resources[0].ID)
	}
}

func TestRetriever_ResourcesForEachDocType(t *testing.T) {
	mockService := &MockReportService{}

	retriever := New(mockService)

	tests := []struct {
		docType            shared.DocType
		expectedResourceID string
	}{
		{shared.DocTypeCPF, "resource-validate-cpf"},
		{shared.DocTypeCNPJ, "resource-check-cnpj"},
		{shared.DocTypePlate, "resource-plate-history"},
		{shared.DocTypeName, "resource-person-search"},
		{shared.DocTypeNone, ""},
	}

	for _, tt := range tests {
		result := retriever.Retrieve("dummy", tt.docType)

		if tt.expectedResourceID == "" {
			if len(result.Resources) != 0 {
				t.Errorf("Expected no resources for DocType %s, got %d", tt.docType, len(result.Resources))
			}
		} else {
			if len(result.Resources) != 1 {
				t.Errorf("Expected 1 resource for DocType %s, got %d", tt.docType, len(result.Resources))
				continue
			}

			if result.Resources[0].ID != tt.expectedResourceID {
				t.Errorf("For DocType %s, expected resource ID '%s', got '%s'",
					tt.docType, tt.expectedResourceID, result.Resources[0].ID)
			}
		}
	}
}
