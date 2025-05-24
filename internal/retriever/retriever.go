package retriever

import (
	"fmt"

	"svindel/internal/shared"
)

type Retriever struct {
	reportService shared.Report
}

func New(reportService shared.Report) *Retriever {
	return &Retriever{
		reportService: reportService,
	}
}

func (r *Retriever) Retrieve(doc string, docType shared.DocType) shared.ContextResult {
	// Fetch reports
	reports, err := r.reportService.GetReportsForDocument(doc, docType)
	if err != nil {
		fmt.Println("Error fetching reports:", err)
		return shared.ContextResult{}
	}

	// Static resource mapping example
	resources := r.resourcesForDocType(docType)

	return shared.ContextResult{
		Reports:   reports,
		Resources: resources,
	}
}

func (r *Retriever) resourcesForDocType(docType shared.DocType) []shared.Resource {
	switch docType {
	case shared.DocTypeCPF:
		return []shared.Resource{
			{
				ID:         "resource-validate-cpf",
				Title:      "Validate CPF",
				HelperText: "Validate a CPF number",
			},
		}
	case shared.DocTypeCNPJ:
		return []shared.Resource{
			{
				ID:         "resource-check-cnpj",
				Title:      "Check CNPJ",
				HelperText: "Check a CNPJ number",
			},
		}
	case shared.DocTypePlate:
		return []shared.Resource{
			{
				ID:         "resource-plate-history",
				Title:      "Plate History",
				HelperText: "Get the history of a plate number",
			},
		}
	case shared.DocTypeName:
		return []shared.Resource{
			{
				ID:         "resource-person-search",
				Title:      "Person Search",
				HelperText: "Search for a person",
			},
		}
	default:
		return []shared.Resource{}
	}
}
