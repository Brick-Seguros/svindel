package loadenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port             string
	OpenaiApiKey     string
	ReportAPIBaseURL string
	ReportAPIToken   string
}

func LoadEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading .env")
		return nil, err
	}

	return &Env{
		Port:             os.Getenv("PORT"),
		OpenaiApiKey:     os.Getenv("OPENAI_API_KEY"),
		ReportAPIBaseURL: os.Getenv("REPORT_API_BASE_URL"),
		ReportAPIToken:   os.Getenv("REPORT_API_TOKEN"),
	}, nil
}
