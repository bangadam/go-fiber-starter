package controller

import (
	"github.com/bangadam/go-fiber-starter/app/module/auth/request"
	"github.com/bangadam/go-fiber-starter/app/module/auth/service"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// do login
// @Summary      Do login
// @Description  API for do login
// @Tags         Authentication
// @Security     Bearer
// @Body 	     request.LoginRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/login [post]
func (_i *authController) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Login(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Login success"},
		Code:     fiber.StatusOK,
	})
}

// Register
// @Summary      Register
// @Description  API for register
// @Tags         Authentication
// @Body 	     request.RegisterRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/register [post]
func (_i *authController) Register(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	res, err := _i.authService.Register(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Data:     res,
		Messages: response.Messages{"Register success"},
		Code:     fiber.StatusOK,
	})
}
