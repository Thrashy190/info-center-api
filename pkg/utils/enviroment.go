package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		Error("Error loading environment")
	} else {
		Succes(".env successfully load")

	}

}
