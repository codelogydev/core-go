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
	c.JSON(200, SuccessResponse{Success: true, Data: data})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{Success: false, Error: message})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, 400, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, 401, message)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, 403, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, 404, message)
}

