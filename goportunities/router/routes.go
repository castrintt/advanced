package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openning", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello workd!",
			})
		})

		v1.GET("/openning/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello workd!",
			})
		})

		v1.POST("/openning", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello workd!",
			})
		})

		v1.PUT("/openning", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello workd!",
			})
		})

		v1.DELETE("/openning", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello workd!",
			})
		})
	}

}
