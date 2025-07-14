package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primaryKey" json:"id"`
	// Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password  string    `json:"-"` // "-" means this won't be included in JSON
	Role      Role      `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
