package handler

import (
	Requests "goportunities/handler/request"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := Requests.CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		return
	}

	entity := request.ToEntity()

	if err := db.Create(&entity).Error; err != nil {
		logger.Errorf("ERROR CREATING OPENING: %v", err)
		return
	}
}
