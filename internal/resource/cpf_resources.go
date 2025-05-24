package resource

import (
	"svindel/internal/shared"
)

var cpfResources = []shared.Resource{
	{
		ID:         "ANALYSIS_PF_BASIC_DATA",
		Title:      "Dados Básicos",
		HelperText: "Consulta os dados cadastrais básicos do CPF.",
	},
	{
		ID:         "ANALYSIS_PF_ESSENTIAL",
		Title:      "Análise Essencial",
		HelperText: "Fornece uma visão essencial dos dados do CPF.",
	},
	{
		ID:         "ANALYSIS_PF_CREDIT_SPC",
		Title:      "Análise de Crédito SPC",
		HelperText: "Verifica restrições e pendências no SPC relacionadas ao CPF.",
	},
	{
		ID:         "ANALYSIS_PF_CREDIT_BOA_VISTA",
		Title:      "Análise de Crédito Boa Vista",
		HelperText: "Consulta informações de crédito na Boa Vista para o CPF.",
	},
	{
		ID:         "ANALYSIS_PF_CREDIT_SERASA",
		Title:      "Análise de Crédito Serasa",
		HelperText: "Verifica restrições de crédito no Serasa para o CPF.",
	},
	{
		ID:         "ANALYSIS_PF_FACIAL_BIOMETRY",
		Title:      "Biometria Facial",
		HelperText: "Realiza a validação biométrica facial da pessoa.",
	},
	{
		ID:         "ANALYSIS_PF_LIVENESS",
		Title:      "Prova de Vida (Liveness)",
		HelperText: "Verifica se a pessoa está viva através de detecção de liveness.",
	},
	{
		ID:         "ANALYSIS_PF_PUBLIC_JOBS",
		Title:      "Cargos Públicos",
		HelperText: "Consulta se o CPF está vinculado a algum cargo público.",
	},
	{
		ID:         "ANALYSIS_PF_SOCIAL_ASSISTANCE",
		Title:      "Assistência Social",
		HelperText: "Verifica se o CPF recebe benefícios sociais ou assistenciais.",
	},
	{
		ID:         "ANALYSIS_PF_CLASS_ENTITIES",
		Title:      "Entidades de Classe",
		HelperText: "Consulta vínculos do CPF com entidades de classe ou associações.",
	},
	{
		ID:         "ANALYSIS_PF_RELATIONS",
		Title:      "Relacionamentos",
		HelperText: "Mapeia os relacionamentos do CPF com outras pessoas ou empresas.",
	},
	{
		ID:         "ANALYSIS_PF_KYC",
		Title:      "KYC (Conheça Seu Cliente)",
		HelperText: "Análise de perfil completa para processos de KYC.",
	},
	{
		ID:         "ANALYSIS_PF_INSS",
		Title:      "Benefícios INSS",
		HelperText: "Consulta vínculos do CPF com benefícios previdenciários (INSS).",
	},
	{
		ID:         "ANALYSIS_PF_CEIS",
		Title:      "CEIS - Cadastro de Servidores",
		HelperText: "Verifica se o CPF consta no Cadastro de Servidores (CEIS).",
	},
	{
		ID:         "ANALYSIS_PF_CNH",
		Title:      "CNH - Carteira Nacional de Habilitação",
		HelperText: "Consulta informações da CNH vinculada ao CPF.",
	},
}
