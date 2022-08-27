package request

import "github.com/bangadam/go-fiber-starter/app/database/schema"

type ArticleRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"required,min=3"`
}

func (req *ArticleRequest) ToDomain() *schema.Article {
	return &schema.Article{
		Title:   req.Title,
		Content: req.Content,
	}
}
