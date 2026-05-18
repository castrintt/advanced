package defaultMessages

import "github.com/gin-gonic/gin"

func SendError(context *gin.Context, statusCode int, message string) {
	context.Header("Content-Type", "application/json")
	context.JSON(statusCode, gin.H{
		"message": message,
	})
}
