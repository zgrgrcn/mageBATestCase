package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func GetEnvVariable(variable string) string {
	getenv := os.Getenv(variable)
	if getenv == "" {
		initEnv()
		getenv = os.Getenv(variable)
	}
	return getenv
}
