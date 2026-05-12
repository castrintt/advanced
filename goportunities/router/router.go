package router

import "github.com/gin-gonic/gin"

const (
	PORT = ":5000"
)

func Initialize() {
	// initialize the router
	router := gin.Default()

	// initialize routes
	routes(router)

	// run the server
	router.Run(PORT)
}
