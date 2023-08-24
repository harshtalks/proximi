package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"gorm.io/gorm"
)

type Businesses struct {
	gorm.Model
	BusinessID  string          `gorm:"uniqueIndex;not null" json:"business_id"`
	Name        string          `json:"name"`
	Address     string          `json:"address"`
	City        string          `json:"city"`
	State       string          `json:"state"`
	PostalCode  string          `json:"postal_code"`
	Latitude    float64         `json:"latitude"`
	Longitude   float64         `json:"longitude"`
	Stars       float64         `json:"stars"`
	ReviewCount int             `json:"review_count"`
	IsOpen      int             `json:"is_open"`
	Attributes  gormjsonb.JSONB `gorm:"type:jsonb" json:"attributes"`
	Categories  string          `json:"categories"`
	Hours       gormjsonb.JSONB `gorm:"type:jsonb" json:"hours"`
	UserID      int             `json:"userID" gorm:"not null"`
	GeoHash     string          `json:"geoHash" gorm:"not null"`
}
