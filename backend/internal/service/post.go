package service

import (
	"fmt"

	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/repository"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type PostService interface {
	CreatePost(userID uuid.UUID, req *models.CreatePostRequest) (*models.Post, error)
	UpdatePost(id uuid.UUID, req *models.UpdatePostRequest) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id uuid.UUID) (*models.Post, error)
	DeletePost(id uuid.UUID) error
}

type postService struct {
	repo       repository.PostRepository
	statusRepo repository.PostStatusRepository
	tagRepo    repository.TagRepository
}

func NewPostService(r repository.PostRepository, sr repository.PostStatusRepository, tr repository.TagRepository) PostService {
	return &postService{repo: r, statusRepo: sr, tagRepo: tr}
}

func (s *postService) CreatePost(userID uuid.UUID, req *models.CreatePostRequest) (*models.Post, error) {
	slug, errSlug := s.generateUniqueSlug(req.Title)
	if errSlug != nil {
		return nil, errSlug
	}

	status, _ := s.getPostStatus(req.StatusValue)

	tags := make([]models.Tag, 0, len(req.Tags))
	for _, name := range req.Tags {
		tag, err := s.tagRepo.FirstOrCreateTagByName(name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}


	post := &models.Post{
		ID:       uuid.New(),
		AuthorID: userID,
		StatusID: status.ID,
		Status:   status,
		Tags:     tags,
		Title:    req.Title,
		Slug:     slug,
		Content:  req.Content,
	}

	err := s.repo.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) UpdatePost(id uuid.UUID, req *models.UpdatePostRequest) (*models.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != post.Title {
		slug, err := s.generateUniqueSlug(req.Title)
		if err != nil {
			return nil, err
		}
		post.Slug = slug
	}

	post.Title = req.Title
	post.Content = req.Content

	status, _ := s.getPostStatus(req.StatusValue)
	post.StatusID = status.ID
	post.Status = status

	tags := make([]models.Tag, 0, len(req.Tags))
	for _, name := range req.Tags {
		tag, err := s.tagRepo.FirstOrCreateTagByName(name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}

	fmt.Println("Tags found or created:", tags)

	post.Tags = tags

	if err := s.repo.Update(post); err != nil {
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

func (s *postService) DeletePost(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *postService) generateUniqueSlug(title string) (string, error) {
	baseSlug := slug.Make(title)
	if baseSlug == "" {
		return "", customErrors.ErrInvalidSlugTitle
	}

	slugCandidate := baseSlug
	suffix := 1

	for {
		exists, err := s.repo.IsSlugExists(slugCandidate)
		if err != nil {
			// Wrap and return repository error
			return "", fmt.Errorf("checking slug existence: %w", err)
		}

		if !exists {
			break
		}

		slugCandidate = fmt.Sprintf("%s-%d", baseSlug, suffix)
		suffix++
		if suffix > 100 {
			return "", customErrors.ErrSlugGenerationExhausted
		}
	}

	return slugCandidate, nil
}

func (s *postService) getPostStatus(statusValue string) (models.PostStatus, error) {
	status, errStatus := s.statusRepo.FindByValue(statusValue)
	if errStatus != nil {
		return models.PostStatus{}, fmt.Errorf("invalid status value: %w", errStatus)
	}

	return status, nil
}
