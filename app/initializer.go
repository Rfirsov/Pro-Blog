package app

import (
	"github.com/Rfirsov/Pro-Blog/database"
	"github.com/Rfirsov/Pro-Blog/internal/user"
)

func InitializeUserService() (user.Repository, user.Service, user.Handler) {
	repo := user.NewRepository(database.DB)
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	return repo, service, handler
}
