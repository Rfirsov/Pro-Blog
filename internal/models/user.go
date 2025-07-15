package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password  string    `json:"-"` // "-" means this won't be included in JSON
	Role      Role      `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User registration structure
type UserRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRegisterSuccessResponse struct {
	Message string `json:"message" example:"User registered successfully"`
	UserID  string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type UserRegisterFailureBadRequestResponse struct {
	Error   string `json:"error" example:"invalid input format"`
	Details string `json:"details,omitempty"`
}

type UserRegisterFailureConflictResponse struct {
	Error string `json:"error" example:"email already registered"`
}

type UserRegisterFailureInternalServerErrorResponse struct {
	Error string `json:"error" example:"password processing failed"`
}

// User login structure
type UserLoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginSuccessResponse struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type UserLoginFailureBadRequestResponse struct {
	Error string `json:"error" example:"invalid login data"`
}

type UserLoginFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

type UserLoginFailureInternalServerErrorResponse struct {
	Error string `json:"error" example:"login process failed"`
}

// Refresh token structure
type UserRefreshTokenSuccessResponse struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	TokenType string `json:"token_type"`
}

type UserRefreshTokenFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"invalid or expired token"`
}

type UserRefreshTokenFailureServerErrorResponse struct {
	Error string `json:"error" example:"token refresh failed"`
}

// Logout user structure
type UserLogoutSuccessResponse struct {
	Message     string `json:"message" example:"Successfully logged out"`
	Instrctions string `json:"instructions" example:"Please remove the token from your client storage"`
}

type UserLogoutFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"invalid or expired token"`
}

// Get user profile structure
type GetUserProfileSuccessResponse struct {
	UserID string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name   string `json:"name" example:"John Doe"`
	Email  string `json:"email" example:"john@example.com"`
	Role   Role   `json:"role" example:"user"`
}

type GetUserProfileFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"user not authenticated"`
}

type GetUserProfileFailureInternalServerErrorResponse struct {
	Error string `json:"error" example:"failed to retrieve user profile"`
}
