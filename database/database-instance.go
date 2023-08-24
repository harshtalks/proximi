package database

import "gorm.io/gorm"

type DatabaseInstance struct {
	DB *gorm.DB
}

// creating a variable for DB struct

var Database DatabaseInstance
