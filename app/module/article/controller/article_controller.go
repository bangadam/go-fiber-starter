package controller

import (
	"strconv"

	"github.com/bangadam/go-fiber-starter/app/module/article/request"
	"github.com/bangadam/go-fiber-starter/app/module/article/service"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	articleService *service.ArticleService
}

//go:generate mockgen -destination=article_controller_mock.go -package=controller . IArticleController
type IArticleController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
}

func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

// get all articles
// @Summary      Get all articles
// @Description  API for getting all articles
// @Tags         Task
// @Security     Bearer
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [get]
func (_i *ArticleController) Index(c *fiber.Ctx) error {
	articles, err := _i.articleService.All()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article list successfully retrieved"},
		Data:     articles,
	})
}

// get one article
// @Summary      Get one article
// @Description  API for getting one article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [get]
func (_i *ArticleController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	articles, err := _i.articleService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article successfully retrieved"},
		Data:     articles,
	})
}

// create article
// @Summary      Create article
// @Description  API for create article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles [post]
func (_i *ArticleController) Store(c *fiber.Ctx) error {
	req := new(request.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	err := _i.articleService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article successfully created"},
	})
}

// update article
// @Summary      update article
// @Description  API for update article
// @Tags         Task
// @Security     Bearer
// @Body 	     request.ArticleRequest
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [put]
func (_i *ArticleController) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	req := new(request.ArticleRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	err = _i.articleService.Update(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article successfully updated"},
	})
}

// delete article
// @Summary      delete article
// @Description  API for delete article
// @Tags         Task
// @Security     Bearer
// @Param        id path int true "Article ID"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /articles/:id [delete]
func (_i *ArticleController) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = _i.articleService.Destroy(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article successfully deleted"},
	})
}
