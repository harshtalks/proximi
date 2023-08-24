package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	// connecting to the remote postgres database instance

	// reading ENVs

	// READING ENVS
	host := os.Getenv("DATABASE_HOST")
	password := os.Getenv("DATABASE_PASSWORD")
	user := os.Getenv("DATABASE_USER")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	// parsing port number for the database

	parsedPort, parsedPortErr := strconv.ParseUint(port, 10, 32)

	if parsedPortErr != nil {
		log.Fatal("Error occured while parsing the port number")
	}

	// creating the database connection string.

	databaseConnectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		user, password, host, parsedPort, name,
	)

	// Postgres Instance

	db, dbErr := gorm.Open(postgres.Open(databaseConnectionString), &gorm.Config{})

	if dbErr != nil {
		log.Fatal("Error Occured while connecting to the database")
	}

	fmt.Println("Connection Opened to Database...")

	Database = DatabaseInstance{
		DB: db,
	}

}
