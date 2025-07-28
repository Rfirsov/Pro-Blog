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
	statusRepo := repository.NewPostStatusRepository(baseRepo)
	tagRepo := repository.NewTagRepository(baseRepo)
	service := service.NewPostService(repo, statusRepo, tagRepo)
	handler := handlers.NewPostHandler(service)
	return handler
}

func InitializePostStatusService() handlers.PostStatusHandler {
	baseRepo := repository.NewBaseRepository(database.DB)

	repo := repository.NewPostStatusRepository(baseRepo)
	service := service.NewPostStatusService(repo)
	handler := handlers.NewPostStatusHandler(service)
	return handler
}
