package service

import (
	"errors"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/middleware"
	"github.com/bangadam/go-fiber-starter/app/module/auth/request"
	"github.com/bangadam/go-fiber-starter/app/module/auth/response"
	user_repo "github.com/bangadam/go-fiber-starter/app/module/user/repository"
)

// AuthService
type articleService struct {
	userRepo user_repo.UserRepository
}

// define interface of IAuthService
//
//go:generate mockgen -destination=article_service_mock.go -package=service . AuthService
type AuthService interface {
	Login(req request.LoginRequest) (res response.LoginResponse, err error)
	Register(req request.RegisterRequest) (res response.RegisterResponse, err error)
}

// init AuthService
func NewAuthService(userRepo user_repo.UserRepository) AuthService {
	return &articleService{
		userRepo: userRepo,
	}
}

func (_i *articleService) Login(req request.LoginRequest) (res response.LoginResponse, err error) {
	// check user by email
	user, err := _i.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("user not found")
		return
	}

	// check password
	if !user.ComparePassword(req.Password) {
		err = errors.New("password not match")
		return
	}

	// do create token
	claims, err := middleware.GenerateTokenAccess(user.ID)
	if err != nil {
		return
	}

	res.Token = claims.Token
	res.Type = claims.Type
	res.ExpiresAt = claims.ExpiresAt.Unix()

	return
}

func (_i *articleService) Register(req request.RegisterRequest) (res response.RegisterResponse, err error) {
	// check user by email
	user, err := _i.userRepo.FindUserByEmail(req.Email)
	if err != nil {
		return
	}

	if user != nil {
		err = errors.New("email already exists")
		return
	}

	// do create user
	newUser := &schema.User{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err = _i.userRepo.CreateUser(newUser)
	if err != nil {
		return
	}

	res.ID = user.ID
	res.Email = user.Email

	return
}
