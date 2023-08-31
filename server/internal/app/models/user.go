package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Address  string
	Timestamp
}

type NewUser struct {
	ID    uuid.UUID
	Name  string
	Email string
}
