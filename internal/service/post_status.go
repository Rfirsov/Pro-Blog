package service

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/repository"
)

type PostStatusService interface {
	GetPostStatuses() ([]models.PostStatus, error)
}

type postStatusService struct {
	repo repository.PostStatusRepository
}

func NewPostStatusService(r repository.PostStatusRepository) PostStatusService {
	return &postStatusService{repo: r}
}

func (s *postStatusService) GetPostStatuses() ([]models.PostStatus, error) {
	return s.repo.FindAll()
}
