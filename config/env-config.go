package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func ConfigureENV(fileName string) {
	// Loading envs

	configureError := godotenv.Load(fileName)

	if configureError != nil {
		fmt.Println("An error occured while loading the env file.")
		log.Fatal("Please provide a valid file name")
	}

	fmt.Println("Successfully loaded env files.....")
}
