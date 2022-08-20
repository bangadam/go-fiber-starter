package controller

import (
	"strconv"

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
// @Router       /artcles [get]
func (_i *ArticleController) Index(c *fiber.Ctx) error {
	articles, err := _i.articleService.GetArticles()
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
// @Router       /artcles/:id [get]
func (_i *ArticleController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	articles, err := _i.articleService.FindOne(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Article list successfully retrieved"},
		Data:     articles,
	})
}
