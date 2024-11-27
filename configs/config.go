package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURI string `json:"db_uri"`
	DBName      string `json:"db_name"`
	Port        string `json:"port"`
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DatabaseURI := os.Getenv("DatabaseURI")
	DBName := os.Getenv("DBName")
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	return &Config{
		DatabaseURI: DatabaseURI,
		DBName:      DBName,
		Port:        portStr,
	}, nil
}
