package models

import (
	"time"
)

type User struct {
	ID           uint
	DisplayName  string
	Username     string
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
