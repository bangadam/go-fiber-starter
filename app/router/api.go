package router

import (
	"github.com/bangadam/go-fiber-starter/storage"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App fiber.Router
}

func NewRouter(fiber *fiber.App) *Router {
	return &Router{
		App: fiber,
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

	// Register routes of modules
	// r.ArticleRouter.RegisterArticleRoutes()
}
