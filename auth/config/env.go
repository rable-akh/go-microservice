package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configD struct {
	DBURI  string
	DBNAME string
}

func config() *configD {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Don't load env file.")
	}
	configs := &configD{
		DBURI:  os.Getenv("DATABASE_URL"),
		DBNAME: os.Getenv("DB_NAME"),
	}

	return configs
}
