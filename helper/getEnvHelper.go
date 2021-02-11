package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVar func
func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file")
	}
	return os.Getenv(key)
}
