package user

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email,omitempty"`
}
