package auth

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) int {
	val, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return val.(int)
}
