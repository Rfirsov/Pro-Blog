package service

import (
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/repository"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type PostService interface {
	CreatePost(userID uuid.UUID, req *models.CreatePostRequest) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id uuid.UUID) (*models.Post, error)
	DeletePost(id uuid.UUID) error
	GenerateSlug(title string) string
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{repo: r}
}

func (s *postService) CreatePost(userID uuid.UUID, req *models.CreatePostRequest) (*models.Post, error) {
	post := &models.Post{
		ID:       uuid.New(),
		AuthorID: userID,
		Title:    req.Title,
		Slug:     slug.Make(req.Title),
		Content:  req.Content,
		Status:   "draft", // Default status
	}

	err := s.repo.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.repo.FindAll()
}

func (s *postService) GetPostByID(id uuid.UUID) (*models.Post, error) {
	return s.repo.FindByID(id)
}

// func (s *postService) UpdatePost(id uuid.UUID, req *models.UpdatePostRequest) (*models.Post, error) {
// }

func (s *postService) DeletePost(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *postService) GenerateSlug(title string) string {
	return slug.Make(title)
}
