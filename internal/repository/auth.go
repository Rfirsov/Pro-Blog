package repository

import (
	customErrors "github.com/Rfirsov/Pro-Blog/internal/errors"
	"github.com/Rfirsov/Pro-Blog/internal/models"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	IfUserExistsByEmail(email string) (bool, error)
	CreateNewUser(newUser *models.User) error
}

type authRepo struct {
	*BaseRepository
}

func NewAuthRepository(base *BaseRepository) AuthRepository {
	return &authRepo{BaseRepository: base}
}

func (r *authRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Select("id", "email", "role", "password").Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *authRepo) IfUserExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.DB.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error

	return exists, err
}

func (r *authRepo) CreateNewUser(newUser *models.User) error {

	// Insert user with transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return customErrors.ErrTransactionStart
	}

	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return customErrors.ErrUserCreation
	}

	if err := tx.Commit().Error; err != nil {
		return customErrors.ErrTransactionCommit
	}

	return nil
}
