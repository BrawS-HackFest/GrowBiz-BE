package utils

import "github.com/gin-gonic/gin"

func HttpInternalError(c *gin.Context, msg string, err error) {
	c.JSON(500, gin.H{
		"status":  "error",
		"message": msg,
		"error":   err,
	})
}

func HttpSuccess(c *gin.Context, msg string, data any) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
		"data":    data,
	})
}

func HttpFailOrError(c *gin.Context, code int, msg string, err error) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": msg,
		"error":   err,
	})
}
