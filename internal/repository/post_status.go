package repository

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
)

type PostStatusRepository interface {
	FindAll() ([]models.PostStatus, error)
	FindByValue(value string) (models.PostStatus, error)
}

type postStatusRepo struct {
	*BaseRepository
}

func NewPostStatusRepository(base *BaseRepository) PostStatusRepository {
	return &postStatusRepo{BaseRepository: base}
}

func (r *postStatusRepo) FindAll() ([]models.PostStatus, error) {
	var statuses []models.PostStatus
	err := r.DB.Find(&statuses).Error
	return statuses, err
}

func (r *postStatusRepo) FindByValue(value string) (models.PostStatus, error) {
	var status models.PostStatus
	err := r.DB.Where("value = ?", value).First(&status).Error

	return status, err
}
