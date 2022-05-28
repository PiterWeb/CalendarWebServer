package crypt

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

func MySecret() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SECRET_CRYPT")
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

