package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// RateLimitMiddleware blocks requests to ensure they adhere to the rate limit
func RateLimitMiddleware(rps int) gin.HandlerFunc {
	rl := ratelimit.New(rps) // per second
	return func(c *gin.Context) {
		rl.Take() // Blocks until the bucket allows
		c.Next()
	}
}
