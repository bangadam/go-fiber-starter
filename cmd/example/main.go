package main

import (
	"go.uber.org/fx"

	"github.com/bangadam/go-fiber-starter/app/middleware"
	"github.com/bangadam/go-fiber-starter/app/module/article"
	"github.com/bangadam/go-fiber-starter/app/router"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap"
	"github.com/bangadam/go-fiber-starter/internal/bootstrap/database"
	"github.com/bangadam/go-fiber-starter/utils/config"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fx.New(
		/* provide patterns */
		// config
		fx.Provide(config.NewConfig),
		// logging
		fx.Provide(bootstrap.NewLogger),
		// fiber
		fx.Provide(bootstrap.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules
		article.NewArticleModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
