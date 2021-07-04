package models

import (
	"time"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	DisplayName  string
	Username     string `gorm:"index:username,unique"`
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
