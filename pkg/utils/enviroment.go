package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironment() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	} else {
		Success(".env successfully load")

	}

}
