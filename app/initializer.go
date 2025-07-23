package app

import (
	"time"

	"github.com/Rfirsov/Pro-Blog/database"
	"github.com/Rfirsov/Pro-Blog/internal/handlers"
	"github.com/Rfirsov/Pro-Blog/internal/repository"
	"github.com/Rfirsov/Pro-Blog/internal/service"
)

func InitializeAuthService() handlers.AuthHandler {
	tokenExpiration := 24 * time.Hour
	baseRepo := repository.NewBaseRepository(database.DB)

	repo := repository.NewAuthRepository(baseRepo)
	service := service.NewAuthService(repo, tokenExpiration)
	handler := handlers.NewAuthHandler(service, tokenExpiration)
	return handler
}

func InitializePostService() handlers.PostHandler {
	baseRepo := repository.NewBaseRepository(database.DB)

	repo := repository.NewPostRepository(baseRepo)
	service := service.NewPostService(repo)
	handler := handlers.NewPostHandler(service)
	return handler
}
