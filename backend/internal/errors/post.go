package customErrors

import "errors"

var (
	ErrFetchPostStatuses       = errors.New("could not fetch post statuses")
	ErrInvalidPostData         = errors.New("invalid post data")
	ErrInvalidSlugTitle        = errors.New("invalid slug title")
	ErrPostCreation            = errors.New("post creation failed")
	ErrPostDelete              = errors.New("post delete failed")
	ErrPostId                  = errors.New("post ID is required")
	ErrPostNotFound            = errors.New("post not found")
	ErrPostStatusNotFound      = errors.New("post status not found")
	ErrPostUpdate              = errors.New("post update failed")
	ErrSlugGenerationExhausted = errors.New("unable to generate a unique slug after many attempts")
)
