package logger

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() func(*gin.Context, io.Writer, time.Duration) zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// https://github.com/gin-contrib/logger/blob/v0.2.2/logger.go#L15
	return func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
		isTerm := isatty.IsTerminal(os.Stdout.Fd())
		logger := zerolog.New(out).
			Output(
				zerolog.ConsoleWriter{
					Out:     out,
					NoColor: !isTerm,
				},
			).
			With().
			Timestamp().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			// Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Logger()

		return logger
	}
}
