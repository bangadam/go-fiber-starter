package service

import (
	"github.com/bangadam/go-fiber-starter/app/module/article/repository"
	"github.com/bangadam/go-fiber-starter/app/module/article/request"
	"github.com/bangadam/go-fiber-starter/app/module/article/response"
)

// ArticleService
type ArticleService struct {
	Repo *repository.ArticleRepository
}

// define interface of ArticleService
type IArticleService interface {
	All() (articles []*response.Article, err error)
	Show(id uint64) (article *response.Article, err error)
	Store(req request.ArticleRequest) (err error)
	Update(id uint64, req request.ArticleRequest) (err error)
	Destroy(id uint64) error
}

// init ArticleService
func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		Repo: repo,
	}
}

// implement interface of ArticleService
func (_i *ArticleService) All() (articles []*response.Article, err error) {
	results, err := _i.Repo.GetArticles()
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		articles = append(articles, response.FromDomain(result))
	}

	return articles, nil
}

func (_i *ArticleService) Show(id uint64) (article *response.Article, err error) {
	result, err := _i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}

	return response.FromDomain(result), nil
}

func (_i *ArticleService) Store(req request.ArticleRequest) (err error) {
	return _i.Repo.Create(req.ToDomain())
}

func (_i *ArticleService) Update(id uint64, req request.ArticleRequest) (err error) {
	return _i.Repo.Update(id, req.ToDomain())
}

func (_i *ArticleService) Destroy(id uint64) error {
	return _i.Repo.Delete(id)
}
