package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/codelogydev/core-go/auth"
	"github.com/codelogydev/core-go/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			response.Unauthorized(c, "missing token")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, err := auth.ValidateToken(token)
		if err != nil {
			response.Unauthorized(c, "invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
