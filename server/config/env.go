package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_, err := os.Stat("./env")
	if os.IsNotExist(err) {
		env, err := godotenv.Unmarshal("AZURE_ORGANIZATION=ORGANIZATION")
		if err != nil {
			log.Fatal(err)
		}
		if err := godotenv.Write(env, "./.env"); err != nil {
			log.Fatal(err)
		}

		return
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
