package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_, err := os.Stat("./.env")
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
