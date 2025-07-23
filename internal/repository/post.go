package repository

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/google/uuid"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindByID(id uuid.UUID) (*models.Post, error)
	FindAll() ([]models.Post, error)
	Update(post *models.Post) error
	Delete(id uuid.UUID) error
	IsSlugExists(slug string) (bool, error)
}

type postRepo struct {
	*BaseRepository
}

func NewPostRepository(base *BaseRepository) PostRepository {
	return &postRepo{BaseRepository: base}
}

func (r *postRepo) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepo) FindByID(id uuid.UUID) (*models.Post, error) {
	var post models.Post
	if err := r.DB.First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepo) FindAll() ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepo) Update(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepo) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Post{}, "id = ?", id).Error
}

func (r *postRepo) IsSlugExists(slug string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Post{}).Where("slug = ?", slug).Count(&count).Error
	return count > 0, err
}
