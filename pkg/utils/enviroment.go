package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnvironment() {

	err := godotenv.Load(".env")
	if err != nil {
		Error("Error loading environment")
	} else {
		Succes(".env successfully load")

	}

}
