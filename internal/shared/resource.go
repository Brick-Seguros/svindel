package shared

type ResourceService interface {
	GetResourcesByDocType(docType DocType) []Resource
}
