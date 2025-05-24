package docext

import (
	shared "svindel/internal/shared"
)

type ContextResult struct {
	Reports   []string // Report IDs
	Resources []string // Resource IDs
}

type Retriever interface {
	Retrieve(doc string, docType shared.DocType) ContextResult
}

type MockRetriever struct{}

func NewRetriever() *MockRetriever {
	return &MockRetriever{}
}

func (r *MockRetriever) Retrieve(doc string, docType shared.DocType) ContextResult {

	switch docType {
	case shared.DocTypeCPF:
		return ContextResult{
			Reports:   []string{"report-cpf-123", "report-cpf-456"},
			Resources: []string{"resource-cpf-validator", "resource-cpf-analyzer"},
		}
	case shared.DocTypeCNPJ:
		return ContextResult{
			Reports:   []string{"report-cnpj-789"},
			Resources: []string{"resource-cnpj-checker"},
		}
	case shared.DocTypePlate:
		return ContextResult{
			Reports:   []string{"report-plate-101"},
			Resources: []string{"resource-plate-verifier"},
		}
	case shared.DocTypeName:
		return ContextResult{
			Reports:   []string{"report-name-001"},
			Resources: []string{"resource-name-search"},
		}
	default:
		return ContextResult{}
	}
}
