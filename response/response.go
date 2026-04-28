package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Success(c *gin.Context, data interface{}) {
	res := SuccessResponse{
		Success: true,
		Data:    data,
	}
	c.JSON(200, res)
}

func Error(c *gin.Context, code int, message string) {
	res := ErrorResponse{
		Success: false,
		Error:   message,
	}
	c.JSON(code, res)
}
