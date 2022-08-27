package response

import (
	"time"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
)

type Article struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(article *schema.Article) (res *Article) {
	if article != nil {
		res = &Article{
			ID: 	article.ID,
			Title:     article.Title,
			Content:   article.Content,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		}
	}

	return res
}
