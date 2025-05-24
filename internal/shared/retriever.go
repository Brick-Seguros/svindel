package shared

type ContextResult struct {
	Reports   []ReportResult
	Resources []Resource
}

type Retriever interface {
	Retrieve(doc string, docType DocType) ContextResult
}
