package controller

import "github.com/bangadam/go-fiber-starter/app/module/article/service"

type Controller struct {
	Article *ArticleController
}

func NewController(articleService *service.ArticleService) *Controller {
	return &Controller{
		Article: &ArticleController{articleService: articleService},
	}
}
