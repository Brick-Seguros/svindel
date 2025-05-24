package resource

import (
	"svindel/internal/shared"
)

var plateResources = []shared.Resource{
	{
		ID:         "ANALYSIS_VEHICLE_APP",
		Title:      "Veículo para APP",
		HelperText: "Verifica se o veículo está registrado para transporte por aplicativos (Uber, 99, etc.).",
	},
	{
		ID:         "ANALYSIS_VEHICLE_TAXI",
		Title:      "Veículo para Táxi",
		HelperText: "Verifica se o veículo está registrado como táxi.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_DEBITS",
		Title:      "Débitos Veiculares",
		HelperText: "Consulta débitos pendentes como IPVA, multas e licenciamento do veículo.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_BASIC_DATA",
		Title:      "Dados Básicos",
		HelperText: "Consulta informações cadastrais básicas do veículo, como marca, modelo, ano e características.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_FULL_REPORT",
		Title:      "Consulta Completa",
		HelperText: "Realiza uma consulta completa com histórico, características, débitos e restrições do veículo.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_AUCTION",
		Title:      "Leilões",
		HelperText: "Verifica se o veículo já passou por leilão.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_GRAVAME",
		Title:      "Gravame",
		HelperText: "Consulta se o veículo possui gravame ativo (restrição financeira).",
	},
	{
		ID:         "ANALYSIS_VEHICLE_SINISTER_INDICATOR",
		Title:      "Indício de Sinistro",
		HelperText: "Indica se há histórico de sinistro grave, perda total ou outros danos severos.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_CAMERA_IMAGES",
		Title:      "Imagens de Câmera (Conector)",
		HelperText: "Consulta imagens recentes do veículo capturadas por câmeras de trânsito (quando disponível).",
	},
	{
		ID:         "ANALYSIS_VEHICLE_BIN_NACIONAL",
		Title:      "BIN Nacional",
		HelperText: "Consulta o histórico do veículo através do BIN Nacional, abrangendo ocorrências e registros em território nacional.",
	},
	{
		ID:         "ANALYSIS_VEHICLE_THEFT_HISTORY",
		Title:      "Histórico de Roubo e Furto",
		HelperText: "Consulta registros de roubo e furto associados à placa do veículo.",
	},
}
