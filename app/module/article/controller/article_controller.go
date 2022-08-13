package controller

import (
	"github.com/bangadam/go-fiber-starter/app/module/article/service"
	"github.com/bangadam/go-fiber-starter/utils/response"
	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	articleService *service.ArticleService
}

type IArticleController interface {
	Index(c *fiber.Ctx) error
}

func NewArticleController(articleService *service.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

// implement interface of IArticleController
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
