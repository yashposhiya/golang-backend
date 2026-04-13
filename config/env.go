package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DB_DSN    string
	Secretkey string
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		Port: os.Getenv("PORT"),
		DB_DSN: os.Getenv("DB_DSN"),
		Secretkey: os.Getenv("SECRET_KEY"),
	}

}
