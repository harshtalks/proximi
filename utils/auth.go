package utils

import (
	"app/database"
	"app/models"
	"errors"

	"gorm.io/gorm"
)

type ErrorUser int

const (
	NOERROR ErrorUser = iota
	NOTFOUND
	CUSTOMERROR
)

// get user by username

func GetUser(username string) (*models.User, ErrorUser) {

	// gorm

	databaseInstance := database.Database.DB

	var user models.User

	if queryError := databaseInstance.Where("username = ?", username).First(&user).Error; queryError != nil {
		if errors.Is(queryError, gorm.ErrRecordNotFound) {
			return nil, NOTFOUND
		}

		return nil, CUSTOMERROR
	}

	return &user, NOERROR
}

// create user

func CreateUser(username string, password string) (*models.User, error) {
	// get the hash of the password

	passwordHash, passwordHashErr := GetPasswordHash(password)

	if passwordHashErr != nil {
		return nil, errors.New("Bad Password")
	}

	var user = models.User{
		Password: passwordHash,
		Username: username,
	}

	databaseInstance := database.Database.DB

	if mutationError := databaseInstance.Create(&user).Error; mutationError != nil {
		return nil, mutationError
	}

	return &user, nil

}
