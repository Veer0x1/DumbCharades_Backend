package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	DB_URL string
	PORT   string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Config = &config{
		DB_URL: os.Getenv("DB_URL"),
		PORT: os.Getenv("PORT"),
	}
}
