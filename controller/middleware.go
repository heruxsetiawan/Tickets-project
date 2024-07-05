package controller

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeDB() {}

// LoggerMiddleware logs all requests and responses
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Process response
		latency := time.Since(startTime)
		status := c.Writer.Status()
		log.Printf("INI LOG BROOO: %s %s %d %s", c.Request.Method, c.Request.URL.Path, status, latency)
	}
}

// JWTMiddleware validates the JWT token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(401, gin.H{"error_1": "Unauthorized_1"})
			c.Abort()
			return
		}

		err := VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error_2": "Unauthorized_2"})
			c.Abort()
			return
		}

		c.Next()
	}
}
