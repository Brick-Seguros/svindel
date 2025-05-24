package main

import (
	"log"

	"svindel/cmd/httpsrv"
	"svindel/internal/chat.go"
	"svindel/internal/completion"
	"svindel/internal/docext"
	"svindel/pkg/loadenv"
)

func main() {

	env, err := loadenv.LoadEnv()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	completion := completion.New(env.OpenaiApiKey)

	extractor := docext.NewExtractor(env.OpenaiApiKey)
	retriever := docext.NewRetriever()
	docExt := docext.New(*extractor, retriever)

	chatService := chat.New(completion, docExt)

	// Setup HTTP server
	router := httpsrv.NewRouter(chatService)
	server := httpsrv.NewServer(env.Port, router)
	server.Start()

}
