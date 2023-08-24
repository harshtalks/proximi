package database

import "app/models"

func Migrate() {
	// migrating the database models here
	println("Migrating the database....")
	Database.DB.AutoMigrate(&models.User{})
	Database.DB.AutoMigrate(&models.Businesses{})
}
