package auth

import (
	"github.com/bangadam/go-fiber-starter/app/module/auth/controller"
	"github.com/bangadam/go-fiber-starter/app/module/auth/service"
	user_repo "github.com/bangadam/go-fiber-starter/app/module/user/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// struct of AuthRouter
type AuthRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

// register bulky of auth module
var NewAuthModule = fx.Options(
	// register repository of auth module
	fx.Provide(user_repo.NewUserRepository),

	// register service of auth module
	fx.Provide(service.NewAuthService),

	// register controller of auth module
	fx.Provide(controller.NewController),

	// register router of auth module
	fx.Provide(NewAuthRouter),
)

// init AuthRouter
func NewAuthRouter(fiber *fiber.App, controller *controller.Controller) *AuthRouter {
	return &AuthRouter{
		App:        fiber,
		Controller: controller,
	}
}

// register routes of auth module
func (_i *AuthRouter) RegisterAuthRoutes() {
	// define controllers
	authController := _i.Controller.Auth

	// define routes
	_i.App.Route("/auth", func(router fiber.Router) {
		router.Post("/login", authController.Login)
		router.Post("/register", authController.Register)
	})
}
