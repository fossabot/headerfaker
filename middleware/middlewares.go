package middleware

import (
	"github.com/gin-gonic/gin"
)

func ServerType() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Server", "HeaderFaker-v0.2.0")
		c.Next()
	}
}
