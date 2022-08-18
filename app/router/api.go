package router

import (
	"github.com/bangadam/go-fiber-starter/app/module/article"
	"github.com/bangadam/go-fiber-starter/docs"
	"github.com/bangadam/go-fiber-starter/storage"
	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	App           fiber.Router
	Cfg           *config.Config
	ArticleRouter *article.ArticleRouter
}

func NewRouter(fiber *fiber.App, cfg *config.Config, articleRouter *article.ArticleRouter) *Router {
	return &Router{
		App:           fiber,
		Cfg:           cfg,
		ArticleRouter: articleRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	r.App.Get("/html", func(c *fiber.Ctx) error {
		example, err := storage.Private.ReadFile("private/example.html")
		if err != nil {
			panic(err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(string(example))
	})

	// Swagger Documentation
	docs.SwaggerInfo.Host = "0.0.0.0:" + r.Cfg.App.Port
	r.App.Get("/swagger/*", swagger.HandlerDefault)

	// Register routes of modules
	r.ArticleRouter.RegisterArticleRoutes()
}
