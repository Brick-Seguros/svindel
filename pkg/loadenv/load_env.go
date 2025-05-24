package loadenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port string

	OpenaiApiKey string
}

func LoadEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading .env")
		return nil, err
	}

	return &Env{
		Port:         os.Getenv("PORT"),
		OpenaiApiKey: os.Getenv("OPENAI_API_KEY"),
	}, nil
}
