package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/codelogydev/core-go/response"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, _ any) {
		response.Error(c, 500, "internal server error")
	})
}
