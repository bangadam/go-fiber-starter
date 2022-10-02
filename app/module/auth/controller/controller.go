package controller

import "github.com/bangadam/go-fiber-starter/app/module/auth/service"

type Controller struct {
	Auth AuthController
}

func NewController(authService service.AuthService) *Controller {
	return &Controller{
		Auth: NewAuthController(authService),
	}
}
