package router

import (
	Handler "goportunities/handler"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {
	// initialize handler
	Handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET(GET_OPENING, Handler.FindOpeningHandler)
		v1.GET(LIST_ALL_OPENING, Handler.FindAllOpeningHandler)
		v1.POST(CREATE_OPENING, Handler.CreateOpeningHandler)
		v1.PUT(UPDATE_OPENING, Handler.UpdateOpeningHandler)
		v1.DELETE(DELETE_OPENING, Handler.DeleteOpeningHandler)
	}
}
