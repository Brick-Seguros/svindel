package resource

import (
	"svindel/internal/shared"
)

var phoneResources = []shared.Resource{
	{
		ID:         "TOOL_PHONE_VALIDATOR",
		Title:      "Validador de Telefone",
		HelperText: "Verifica se o número de telefone é válido, ativo e se está corretamente formatado.",
	},
	{
		ID:         "TOOL_PHONE_IMEI_CHECK",
		Title:      "Consulta de IMEI",
		HelperText: "Consulta o IMEI do dispositivo associado ao telefone para verificar restrições, bloqueios ou irregularidades.",
	},
}
