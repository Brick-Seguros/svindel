package resource

import (
	"svindel/internal/shared"
)

type ResourceService struct {
	resources map[shared.DocType][]shared.Resource
}

func New() *ResourceService {
	return &ResourceService{
		resources: map[shared.DocType][]shared.Resource{
			shared.DocTypeCPF:     cpfResources,
			shared.DocTypeCNPJ:    cnpjResources,
			shared.DocTypePlate:   plateResources,
			shared.DocTypeName:    nameResources,
			shared.DocTypeEmail:   emailResources,
			shared.DocTypePhone:   phoneResources,
			shared.DocTypeAddress: addressesResources,
		},
	}
}

func (s *ResourceService) GetResourcesByDocType(docType shared.DocType) []shared.Resource {
	return s.resources[docType]
}
