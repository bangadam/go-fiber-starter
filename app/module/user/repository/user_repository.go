package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
)

type userRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . UserRepository
type UserRepository interface {
	FindUserByUsername(username string) (user *schema.User, err error)
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// implement interface of IUserRepository
func (_i *userRepository) FindUserByUsername(username string) (user *schema.User, err error) {
	if err := _i.DB.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
