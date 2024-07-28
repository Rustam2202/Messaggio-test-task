package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func NewLogger(level string) {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
