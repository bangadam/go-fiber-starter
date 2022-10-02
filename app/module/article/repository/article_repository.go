package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/module/article/request"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/paginator"
)

type articleRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . ArticleRepository
type ArticleRepository interface {
	GetArticles(req request.ArticlesRequest) (articles []*schema.Article, paging paginator.Pagination, err error)
	FindOne(id uint64) (article *schema.Article, err error)
	Create(article *schema.Article) (err error)
	Update(id uint64, article *schema.Article) (err error)
	Delete(id uint64) (err error)
}

func NewArticleRepository(db *database.Database) ArticleRepository {
	return &articleRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *articleRepository) GetArticles(req request.ArticlesRequest) (articles []*schema.Article, paging paginator.Pagination, err error) {
	var count int64

	query := _i.DB.DB.Model(&schema.Article{})
	query.Count(&count)

	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)

	err = query.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&articles).Error
	if err != nil {
		return
	}

	paging = *req.Pagination

	return
}

func (_i *articleRepository) FindOne(id uint64) (article *schema.Article, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *articleRepository) Create(article *schema.Article) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *articleRepository) Update(id uint64, article *schema.Article) (err error) {
	return _i.DB.DB.Model(&schema.Article{}).
		Where(&schema.Article{ID: id}).
		Updates(article).Error
}

func (_i *articleRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Article{}, id).Error
}
