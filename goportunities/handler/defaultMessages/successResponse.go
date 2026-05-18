package defaultMessages

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(context *gin.Context, operation string, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"operation": fmt.Sprintf("%s operation successfully", operation),
		"data":      data,
	})
}
