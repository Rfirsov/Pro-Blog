package repository

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/google/uuid"
)

type PostRepository interface {
	Create(post *models.Post) error
	Update(post *models.Post) error
	FindByID(id uuid.UUID) (*models.Post, error)
	FindAll() ([]models.Post, error)
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

func (r *postRepo) Update(updatedPost *models.Post) error {
	err := r.DB.Model(updatedPost).Where("id = ?", updatedPost.ID).Updates(map[string]interface{}{
		"title":     updatedPost.Title,
		"author_id": updatedPost.AuthorID,
		"status_id": updatedPost.StatusID,
		"slug":      updatedPost.Slug,
		"content":   updatedPost.Content,
	}).Error

	return err
}

func (r *postRepo) FindByID(id uuid.UUID) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Status").First(&post, "id = ?", id).Error; err != nil {
		return nil, err
	}

	post.StatusValue = post.Status.Value

	return &post, nil
}

func (r *postRepo) FindAll() ([]models.Post, error) {
	var posts []models.Post
	if err := r.DB.Preload("Status").Find(&posts).Error; err != nil {
		return nil, err
	}

	for i := range posts {
		posts[i].StatusValue = posts[i].Status.Value
	}

	return posts, nil
}

func (r *postRepo) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Post{}, "id = ?", id).Error
}

func (r *postRepo) IsSlugExists(slug string) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Post{}).Where("slug = ?", slug).Count(&count).Error
	return count > 0, err
}
