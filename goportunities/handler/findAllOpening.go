package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}
