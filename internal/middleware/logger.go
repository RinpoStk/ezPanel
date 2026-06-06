package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rinpostk/ezPanel/internal/logger"
)

// Logger returns a Gin middleware that logs request method, path, status and latency.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("[%d] %s %s %v", status, c.Request.Method, path, latency)
	}
}
