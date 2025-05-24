package resource

import (
	"svindel/internal/shared"
)

var cnpjResources = []shared.Resource{
	{
		ID:         "ANALYSIS_PJ_ESSENTIAL",
		Title:      "Análise Essencial PJ",
		HelperText: "Fornece uma visão essencial dos dados cadastrais da empresa.",
	},
	{
		ID:         "ANALYSIS_PJ_CREDIT_SPC",
		Title:      "Análise de Crédito SPC PJ",
		HelperText: "Verifica restrições e pendências no SPC relacionadas ao CNPJ.",
	},
	{
		ID:         "ANALYSIS_PJ_CREDIT_BOA_VISTA",
		Title:      "Análise de Crédito Boa Vista PJ",
		HelperText: "Consulta informações de crédito na Boa Vista para a empresa.",
	},
	{
		ID:         "ANALYSIS_PJ_CREDIT_SERASA",
		Title:      "Análise de Crédito Serasa PJ",
		HelperText: "Verifica restrições de crédito no Serasa vinculadas ao CNPJ.",
	},
	{
		ID:         "ANALYSIS_PJ_INCOME_STATEMENT",
		Title:      "Demonstrativo de Resultados (DRE)",
		HelperText: "Consulta o demonstrativo de resultados da empresa (DRE).",
	},
	{
		ID:         "ANALYSIS_PJ_BALANCE_SHEET",
		Title:      "Balanço Patrimonial",
		HelperText: "Consulta o balanço patrimonial da empresa.",
	},
	{
		ID:         "ANALYSIS_PJ_KYB",
		Title:      "KYB (Conheça Sua Empresa)",
		HelperText: "Análise completa para processos de KYB, validando dados da empresa.",
	},
	{
		ID:         "ANALYSIS_PJ_SINTEGRA",
		Title:      "Consulta Sintegra",
		HelperText: "Consulta informações fiscais da empresa diretamente no Sintegra.",
	},
}
