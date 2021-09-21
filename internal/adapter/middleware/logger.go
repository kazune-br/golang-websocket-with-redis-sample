package middleware

import (
	"io"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetLogger() gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithUTC(true),
		logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
			return zerolog.New(out).With().
				// default key values
				Timestamp().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user_agent", c.Request.UserAgent()).
				// custom key values
				Str("tag", "access").
				Logger()
		}))
}
