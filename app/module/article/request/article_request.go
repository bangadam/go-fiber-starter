package request

type ArticleRequest struct {
	Title string `json:"title" form:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=3"`
}