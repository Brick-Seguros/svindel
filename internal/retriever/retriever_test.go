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

type MockResourceService struct {
	ExpectedResources map[shared.DocType][]shared.Resource
}

func (m *MockResourceService) GetResourcesByDocType(docType shared.DocType) []shared.Resource {
	return m.ExpectedResources[docType]
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

	mockResources := []shared.Resource{
		{
			ID:         "resource-validate-cpf",
			Title:      "Fraud Check",
			HelperText: "Check if the document is fraudulent",
		},
	}

	mockResourceService := &MockResourceService{
		ExpectedResources: map[shared.DocType][]shared.Resource{
			shared.DocTypeCPF: mockResources,
		},
	}

	retriever := New(mockService, mockResourceService)

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
	mockResourceService := &MockResourceService{
		ExpectedResources: map[shared.DocType][]shared.Resource{
			shared.DocTypeCPF: {
				{
					ID:         "resource-validate-cpf",
					Title:      "Validate CPF",
					HelperText: "Check if the document is valid",
				},
			},
			shared.DocTypeCNPJ: {
				{
					ID:         "resource-check-cnpj",
					Title:      "Check CNPJ",
					HelperText: "Check if the document is valid",
				},
			},
			shared.DocTypePlate: {
				{
					ID:         "resource-plate-history",
					Title:      "Plate History",
					HelperText: "Check if the plate is valid",
				},
			},
			shared.DocTypeName: {
				{
					ID:         "resource-person-search",
					Title:      "Person Search",
					HelperText: "Check if the person is valid",
				},
			},
		},
	}

	retriever := New(mockService, mockResourceService)

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
