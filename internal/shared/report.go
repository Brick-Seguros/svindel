package shared

type Report interface {
	GetReportsForDocument(doc string, docType DocType) ([]ReportResult, error)
}

type ReportResult struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Document  string `json:"document"`
	CreatedAt string `json:"created_at"`
}
