package user

import (
	"errors"
)

type Service interface {
	CreateUser(req CreateUserRequest) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateUser(req CreateUserRequest) (*User, error) {
	if existing, _ := s.repo.GetByEmail(req.Email); existing.ID != [16]byte{} {
		return nil, errors.New("email already registered")
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.CreateUser(user)
	return user, err
}
