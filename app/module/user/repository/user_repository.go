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
	FindUserByEmail(email string) (user *schema.User, err error)
	CreateUser(user *schema.User) (res *schema.User, err error)
}

func NewUserRepository(db *database.Database) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (_i *userRepository) FindUserByEmail(email string) (user *schema.User, err error) {
	if err := _i.DB.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (_i *userRepository) CreateUser(user *schema.User) (res *schema.User, err error) {
	if err := _i.DB.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
