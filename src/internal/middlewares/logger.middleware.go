package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate elapsed time
		latency := time.Since(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Log request details
		log.Printf("| %d | %s | %s | %s | %s | %s |",
			statusCode,
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			latency,
			c.Request.UserAgent(),
		)
	}
}
