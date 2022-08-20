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
	FindOne(id int) (*ent.Article, error)
	Create(article *ent.Article) (*ent.Article, error)
	Update(id int, article *ent.Article) (*ent.Article, error)
	Delete(id int) error
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

func (_i *ArticleRepository) FindOne(id int) (*ent.Article, error) {
	return _i.DB.Ent.Article.Query().Where(article.IDEQ(id)).First(context.Background())
}

func (_i *ArticleRepository) Create(article *ent.Article) (*ent.Article, error) {
	return _i.DB.Ent.Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(context.Background())
}

func (_i *ArticleRepository) Update(id int, article *ent.Article) (*ent.Article, error) {
	return _i.DB.Ent.Article.UpdateOneID(id).
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(context.Background())
}

func (_i *ArticleRepository) Delete(id int) error {
	return _i.DB.Ent.Article.DeleteOneID(id).Exec(context.Background())
}