package shared

type DocType string

const (
	DocTypeCPF     DocType = "CPF"
	DocTypeCNPJ    DocType = "CNPJ"
	DocTypeName    DocType = "NAME"
	DocTypeEmail   DocType = "EMAIL"
	DocTypePhone   DocType = "PHONE"
	DocTypeAddress DocType = "ADDRESS"
	DocTypePlate   DocType = "PLATE"
	DocTypeNone    DocType = "NONE" // No document found
)

type ExtractionResult struct {
	Document     string
	DocumentType DocType
	IsQuestion   bool
}
