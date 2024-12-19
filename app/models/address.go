package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID uint
	UserID string `gorm:"not null"`
	StreetAddress   string `gorm:"not null"`
	City string `gorm:"not null"`
	State string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	IsPrimary bool `gorm:"not null;default:false"`
}
