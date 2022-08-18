package repository

import (
	"context"

	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/internal/ent"
	article "github.com/bangadam/go-fiber-starter/internal/ent/article"
)

type ArticleRepository struct {
	DB *database.Database
}

//go:generate mockgen -destination=article_repository_mock.go -package=repository . IArticleRepository
type IArticleRepository interface {
	GetArticles() ([]*ent.Article, error)
}

func NewArticleRepository(db *database.Database) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

// implement interface of IArticleRepository
func (_i *ArticleRepository) GetArticles() ([]*ent.Article, error) {
	return _i.DB.Ent.Article.Query().Order(ent.Asc(article.FieldID)).All(context.Background())
}
