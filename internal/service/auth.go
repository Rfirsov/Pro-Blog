package service

import (
	"regexp"
	"time"

	"github.com/Rfirsov/Pro-Blog/config"
	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/models"
	"github.com/Rfirsov/Pro-Blog/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	IfUserExists(email string) (bool, error)
	Registration(newUser *models.User) error
	GetUser(req models.UserLoginRequest) (*models.User, error)
	GenerateJWT(user *models.User) (string, error)
	GenerateRefreshJWT(userID uuid.UUID) (string, error)
	ValidateEmailFormat(user *models.UserRegisterRequest) error
}

type authService struct {
	repo            repository.AuthRepository
	jwtSecret       []byte
	tokenExpiration time.Duration
}

func NewAuthService(r repository.AuthRepository, tokenExpiration time.Duration) AuthService {
	return &authService{repo: r, jwtSecret: []byte(config.Cfg.JWT.Secret), tokenExpiration: tokenExpiration}
}

func (s *authService) IfUserExists(email string) (bool, error) {
	exists, err := s.repo.IfUserExistsByEmail(email)
	return exists, err
}

// Register handles user registration
func (s *authService) Registration(newUser *models.User) error {
	err := s.repo.CreateNewUser(newUser)
	return err

}

func (s *authService) GetUser(req models.UserLoginRequest) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	return user, err
}

func (s *authService) generateTokenWithClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)

	return tokenString, err
}

func (s *authService) GenerateJWT(user *models.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"email":   user.Email,
		"iat":     now.Unix(),
		"exp":     now.Add(s.tokenExpiration).Unix(),
	}

	tokenString, err := s.generateTokenWithClaims(claims)

	return tokenString, err
}

func (s *authService) GenerateRefreshJWT(userID uuid.UUID) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"iat":     now.Unix(),
		"exp":     now.Add(s.tokenExpiration).Unix(),
	}

	tokenString, err := s.generateTokenWithClaims(claims)

	return tokenString, err
}

// Validate checks if email format is valid
func (s *authService) ValidateEmailFormat(u *models.UserRegisterRequest) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return customErrors.ErrInvalidEmailFormat
	}
	return nil
}
