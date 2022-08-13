package bootstrap

import (
	"os"

	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// initialize logger
func NewLogger(cfg *config.Config) zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(cfg.Logger.Level)

	return log.Hook(PreforkHook{})
}

// prefer hook for zerologger
type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}
