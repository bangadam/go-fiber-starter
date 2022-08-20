package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/article/repository"
	"github.com/bangadam/go-fiber-starter/app/module/article/request"
	"github.com/bangadam/go-fiber-starter/internal/ent"
)

// ArticleService
type ArticleService struct {
	Repo *repository.ArticleRepository
}

// define interface of ArticleService
type IArticleService interface {
	All() ([]*ent.Article, error)
	Show(id int) (*ent.Article, error)
	Store(req request.ArticleRequest) (*ent.Article, error)
	Update(id int, req request.ArticleRequest) (*ent.Article, error)
	Destroy(id int) error
}

// init ArticleService
func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *ArticleService) All() ([]*ent.Article, error) {
	return _i.Repo.GetArticles()
}

func (_i *ArticleService) Show(id int) (*ent.Article, error) {
	return _i.Repo.FindOne(id)
}

func (_i *ArticleService) Store(req request.ArticleRequest) (*ent.Article, error) {
	return _i.Repo.Create(&ent.Article{
		Title:   req.Title,
		Content: req.Content,
	})
}

func (_i *ArticleService) Update(id int, req request.ArticleRequest) (*ent.Article, error) {
	return _i.Repo.Update(id, &ent.Article{
		Title:   req.Title,
		Content: req.Content,
	})
}

func (_i *ArticleService) Destroy(id int) error {
	return _i.Repo.Delete(id)
}