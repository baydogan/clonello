package services

import (
	"errors"
	"github.com/baydogan/clonello/user-service/internal/models"
	"github.com/baydogan/clonello/user-service/internal/utils"
	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user with the same username already exists")

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(user *models.User, password string) error {
	var existingUser models.User

	if err := s.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword

	return s.DB.Create(user).Error
}

func (s *UserService) LoginUser(username, password string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	if err := utils.CheckPassword(user.HashedPassword, password); err != nil {
		return nil, err
	}
	return &user, nil
}
