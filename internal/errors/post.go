package customErrors

import "errors"

var (
	ErrInvalidPostData         = errors.New("invalid post data")
	ErrPostCreation            = errors.New("post creation failed")
	ErrPostDelete              = errors.New("post delete failed")
	ErrPostId                  = errors.New("post ID is required")
	ErrPostNotFound            = errors.New("post not found")
	ErrPostUpdate              = errors.New("post update failed")
	ErrInvalidSlugTitle        = errors.New("invalid slug title")
	ErrSlugGenerationExhausted = errors.New("unable to generate a unique slug after many attempts")
)
