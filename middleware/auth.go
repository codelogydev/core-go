package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/codelogydev/core-go/auth"
)

type errorBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			c.JSON(http.StatusUnauthorized, errorBody{Success: false, Error: "missing token"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		userID, err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errorBody{Success: false, Error: "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
