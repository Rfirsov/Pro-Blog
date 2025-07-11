package models

import (
	"time"
)

type User struct {
	ID int `json:"id"`
	// Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password  string    `json:"-"` // "-" means this won't be included in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserLogin represents login request data
type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserRegister represents registration request data
type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
