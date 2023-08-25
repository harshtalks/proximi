package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func ConfigureENV() {
	// Loading envs

	configureError := godotenv.Load()

	if configureError != nil {
		fmt.Println("An error occured while loading the env file.")
		log.Fatal("Please provide a valid file name")
	}

	fmt.Println("Successfully loaded env files.....")
}
