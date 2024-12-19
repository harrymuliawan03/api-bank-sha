package models

import (
	"gorm.io/gorm"
)

type ConnectApps struct {
	gorm.Model
	ID uint
	UserID string `gorm:"not null"`
	AppName   string `gorm:"not null"`
}
