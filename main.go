package main

import (
	"log"
	"net/http"

	"svindel/cmd/httpsrv"
	"svindel/internal/chat"
	"svindel/internal/completion"
	"svindel/internal/docext"
	"svindel/internal/report"
	reportinfra "svindel/internal/report/infra"
	"svindel/internal/resource"
	"svindel/internal/retriever"
	"svindel/pkg/loadenv"
)

func main() {

	env, err := loadenv.LoadEnv()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	reportRepo := reportinfra.NewReportAPIRepository(env.ReportAPIBaseURL, env.ReportAPIToken, &http.Client{})
	reportService := report.New(reportRepo)

	resourceService := resource.New()

	completion := completion.New(env.OpenaiApiKey)

	extractor := docext.NewOpenAiExtractor(env.OpenaiApiKey)
	retriever := retriever.New(reportService, resourceService)

	docExt := docext.New(extractor, retriever)

	chatService := chat.New(completion, docExt)

	// Setup HTTP server
	router := httpsrv.NewRouter(chatService)
	server := httpsrv.NewServer(env.Port, router)
	server.Start()

}
