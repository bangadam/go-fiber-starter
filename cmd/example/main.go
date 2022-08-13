package main

import (
	"go.uber.org/fx"

	"github.com/bangadam/go-fiber-starter/internal/bootstrap"
	"github.com/bangadam/go-fiber-starter/utils/config"
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
	// middleware
	// router

	// provide modules
	// article

	// start aplication
	// call bootstrap

	// define logger

	).Run()
}
