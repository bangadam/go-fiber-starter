package repository

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
)

type ArticleRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . IArticleRepository
type IArticleRepository interface {
	GetArticles() (articles []*schema.Article, err error)
	FindOne(id uint64) (article *schema.Article, err error)
	Create(article *schema.Article) (err error)
	Update(id uint64, article *schema.Article) (err error)
	Delete(id uint64) (err error)
}

func NewArticleRepository(db *database.Database) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *ArticleRepository) GetArticles() (articles []*schema.Article, err error) {
	if err := _i.DB.DB.Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (_i *ArticleRepository) FindOne(id uint64) (article *schema.Article, err error) {
	if err := _i.DB.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (_i *ArticleRepository) Create(article *schema.Article) (err error) {
	return _i.DB.DB.Create(article).Error
}

func (_i *ArticleRepository) Update(id uint64, article *schema.Article) (err error) {
	return _i.DB.DB.Model(&schema.Article{}).
		Where(&schema.Article{ID: id}).
		Updates(article).Error
}

func (_i *ArticleRepository) Delete(id uint64) error {
	return _i.DB.DB.Delete(&schema.Article{}, id).Error
}
