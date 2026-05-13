package handler

import (
	Requests "goportunities/handler/requests"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := Requests.CreateOpeningRequest{}
	context.BindJSON(&request)
	logger.Infof("Request: %+v", request)
}
