package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBIP       string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("error in loading .env file")
	}

	return &Config{
		DBIP:       os.Getenv("DB_IP"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}
