package middleware

import (
	"github.com/gin-gonic/gin"
)

func ServerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		serverName := "UIT - 2024"
		c.Writer.Header().Set("Server", serverName)

	}
}
