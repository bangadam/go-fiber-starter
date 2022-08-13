package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/article/repository"
	"github.com/bangadam/go-fiber-starter/internal/ent"
)

// ArticleService
type ArticleService struct {
	Repo *repository.ArticleRepository
}

// define interface of ArticleService
type IArticleService interface {
	GetArticles() ([]*ent.Article, error)
}

// init ArticleService
func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *ArticleService) GetArticles() ([]*ent.Article, error) {
	return _i.Repo.GetArticles()
}
