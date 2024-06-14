package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Cannot load .env file:%v", err.Error())
	}
}

func ReadOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Cannot load environment variable `%v` from file .env file", key)
	}
	return value
}
