package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}

func ListAllOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello workd!",
	})
}
