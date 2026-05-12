package router

import "github.com/gin-gonic/gin"

const (
	PORT = ":5000"
)

func Initialize() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	router.Run(PORT)
}
