package repository

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
)

type TagRepository interface {
	FirstOrCreateTagByName(name string) (*models.Tag, error)
}

type tagRepo struct {
	*BaseRepository
}

func NewTagRepository(base *BaseRepository) TagRepository {
	return &tagRepo{BaseRepository: base}
}

func (r *tagRepo) FirstOrCreateTagByName(name string) (*models.Tag, error) {
	tag := &models.Tag{}
	result := r.DB.Where("name = ?", name).FirstOrCreate(tag, models.Tag{Name: name})
	if result.Error != nil {
		return nil, result.Error
	}
	return tag, nil
}
