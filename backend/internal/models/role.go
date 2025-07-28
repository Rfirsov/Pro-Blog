package models

type Role string

const (
	RoleUser   Role = "user"
	RoleAuthor Role = "author"
	RoleAdmin  Role = "admin"
)
