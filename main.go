package main

import (
	"log"
	"net/http"

	"svindel/cmd/httpsrv"
	"svindel/internal/chat"
	"svindel/internal/completion"
	"svindel/internal/docext"
	"svindel/internal/eval"
	evalinfra "svindel/internal/eval/infra"
	entitystrat "svindel/internal/eval/strategies/entity"
	hallucinationstrat "svindel/internal/eval/strategies/hallucination"
	relevancestrat "svindel/internal/eval/strategies/relevance"
	"svindel/internal/report"
	reportinfra "svindel/internal/report/infra"
	"svindel/internal/resource"
	"svindel/internal/retriever"
	"svindel/internal/shared"
	"svindel/pkg/loadenv"
)

func main() {

	env, err := loadenv.LoadEnv()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	entityEvalStrategy := entitystrat.New(env.OpenaiApiKey)
	hallucinationEvalStrategy := hallucinationstrat.New(env.OpenaiApiKey)
	relevanceEvalStrategy := relevancestrat.New(env.OpenaiApiKey)

	evalRepo := evalinfra.New()
	evalService := eval.New([]shared.Strategy{entityEvalStrategy, hallucinationEvalStrategy, relevanceEvalStrategy}, evalRepo)

	reportRepo := reportinfra.NewReportAPIRepository(env.ReportAPIBaseURL, env.ReportAPIToken, &http.Client{})
	reportService := report.New(reportRepo)

	resourceService := resource.New()

	completion := completion.New(env.OpenaiApiKey)

	extractor := docext.NewOpenAiExtractor(env.OpenaiApiKey, evalService)
	retriever := retriever.New(reportService, resourceService)

	docExt := docext.New(extractor, retriever)

	chatService := chat.New(completion, docExt)

	// Setup HTTP server
	router := httpsrv.NewRouter(chatService)
	server := httpsrv.NewServer(env.Port, router)
	server.Start()

}
