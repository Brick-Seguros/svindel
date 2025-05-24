package shared

type DocType string

const (
	DocTypeCPF   DocType = "CPF"
	DocTypeCNPJ  DocType = "CNPJ"
	DocTypeName  DocType = "NAME"
	DocTypePlate DocType = "PLATE"
	DocTypeNone  DocType = "NONE" // No document found
)

type ExtractionResult struct {
	Document     string
	DocumentType DocType
	IsQuestion   bool
}
