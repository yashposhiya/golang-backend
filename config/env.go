package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DB_DSN    string
	Secretkey string
	JWT_XPR   int
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExp, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		log.Fatal("Invalid JWT_EXP value")
	}

	AppConfig = Config{
		Port:      os.Getenv("PORT"),
		DB_DSN:    os.Getenv("DB_DSN"),
		Secretkey: os.Getenv("SECRET_KEY"),
		JWT_XPR:   jwtExp,
	}
}
