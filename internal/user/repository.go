package repository

import (
	"github.com/Rfirsov/Pro-Blog/models"
	"github.com/Rfirsov/Pro-Blog/database"
)

func CreateUser(u *user.User) error {
	return database.DB.Create(u).Error
}
