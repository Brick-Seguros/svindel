package resource

import (
	"svindel/internal/shared"
)

var addressesResources = []shared.Resource{
	{
		ID:         "TOOL_ADDRESS_VALIDATOR",
		Title:      "Validador de Endereço",
		HelperText: "Verifica se o endereço é válido, existe e está corretamente cadastrado.",
	},
	{
		ID:         "TOOL_GOOGLE_SEARCH",
		Title:      "Busca Google",
		HelperText: "Realiza uma busca no Google com base no endereço para verificar informações públicas e referências.",
	},
	{
		ID:         "TOOL_RISK_AREA",
		Title:      "Consulta de Área de Risco",
		HelperText: "Verifica se o endereço está localizado em áreas de risco, como zonas com alta criminalidade ou restrições de entrega.",
	},
	{
		ID:         "TOOL_CORREIOS_DELIVERY_AREA",
		Title:      "Região de Entrega dos Correios (Risco)",
		HelperText: "Verifica se o endereço está em regiões com restrições ou riscos de entrega, segundo os Correios.",
	},
}
