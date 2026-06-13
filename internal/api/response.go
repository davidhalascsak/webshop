package api

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, APIResponse{
		Code: 200,
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, err error) {
	c.JSON(code, APIResponse{
		Code:    code,
		Message: message,
		Error:   err.Error(),
	})
}
