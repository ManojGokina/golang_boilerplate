package middleware

import (
	"net/http"
	"time"

	"backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Logger middleware for request logging
func Logger(log logger.Logger) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log.Info("Requesst: ",
			"method=", param.Method,
			" path=", param.Path,
			" status=", param.StatusCode,
			" latency=", param.Latency,
			" ip=", param.ClientIP,
			" user-agent=", param.Request.UserAgent(),
		)
		return ""
	})
}

// Recovery middleware for panic recovery
func Recovery(log logger.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		log.Error("Panic recovered: ", recovered)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": "Something went wrong",
		})
	})
}

// CORS middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RateLimit middleware
func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second), 100) // 100 requests per second

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Rate limit exceeded",
				"message": "Too many requests",
			})
			return
		}
		c.Next()
	}
}

// Auth middleware (placeholder for JWT authentication)
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT token validation
		// token := c.GetHeader("Authorization")
		// if token == "" {
		//     c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		//     return
		// }
		c.Next()
	}
}