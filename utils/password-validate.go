package utils

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const minEntropy = 60

func IsPasswordStrong(password string) error {
	return passwordvalidator.Validate(password, minEntropy)
}

func VerifyPasswordHash(hashedPassword string, password string) bool {
	hashedPasswordByte := []byte(hashedPassword)
	passwordByte := []byte(password)

	if bcryptError := bcrypt.CompareHashAndPassword(hashedPasswordByte, passwordByte); bcryptError != nil {
		return false
	}

	return true
}

func GetPasswordHash(password string) (string, error) {
	passwordByte := []byte(password)

	hashedPassword, bcryptErr := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)

	return string(hashedPassword), bcryptErr
}
