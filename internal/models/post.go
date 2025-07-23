package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug      string    `gorm:"uniqueIndex;not null" json:"slug"`
	Content   string    `gorm:"type:text" json:"content"`
	Status    string    `gorm:"default:'draft'" json:"status"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Post creation structure
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=255"`
	Content string `json:"content" binding:"required,min=10"`
}

type CreatePostSuccessResponse struct {
	Post Post `json:"post"`
}

type CreatePostFailureBadRequestResponse struct {
	Error string `json:"error" example:"invalid post data"`
}

type CreatePostFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

type CreatePostFailureInternalServerErrorResponse struct {
	Error   string `json:"error" example:"post creation failed"`
	Details string `json:"details" example:"error details if available"`
}

// Post update structure
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"min=3,max=255"`
	Content string `json:"content" binding:"min=10"`
}

type UpdatePostSuccessResponse struct {
	Message string `json:"message" example:"post updated successfully"`
	Post    Post   `json:"post"`
}

type UpdatePostFailureBadRequestResponse struct {
	Error string `json:"error" example:"invalid post data"`
}

type UpdatePostFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"invalid credentials"`
}

type UpdatePostFailureInternalServerErrorResponse struct {
	Error   string `json:"error" example:"post update failed"`
	Details string `json:"details" example:"error details if available"`
}

// Get all posts response structure
type GetAllPostsSuccessResponse struct {
	Posts []Post `json:"posts"`
}

type GetAllPostsFailureInternalServerErrorResponse struct {
	Error   string `json:"error" example:"error retrieving posts"`
	Details string `json:"details" example:"error details if available"`
}

// Post retrieval structure
type GetPostByIDSuccessResponse struct {
	Post Post `json:"post"`
}

type GetPostByIDFailureBadRequestResponse struct {
	Error string `json:"error" example:"post not found"`
}

type GetPostByIDFailureInternalServerErrorResponse struct {
	Error   string `json:"error" example:"error retrieving post"`
	Details string `json:"details" example:"error details if available"`
}

// Post delete structure
type DeletePostSuccessResponse struct {
	Message string    `json:"message" example:"post deleted successfully"`
	ID      uuid.UUID `json:"post_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type DeletePostFailureBadRequestResponse struct {
	Error string `json:"error" example:"post not found"`
}

type DeletePostFailureUnauthorizedResponse struct {
	Error string `json:"error" example:"authorization header missing"`
}

type DeletePostFailureInternalServerErrorResponse struct {
	Error   string `json:"error" example:"post delete failed"`
	Details string `json:"details" example:"error details if available"`
}
