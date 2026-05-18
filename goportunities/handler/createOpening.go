package handler

import (
	"goportunities/handler/defaultMessages"
	requests "goportunities/handler/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := requests.CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		defaultMessages.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	entity := request.ToEntity()

	if err := db.Create(&entity).Error; err != nil {
		logger.Errorf("ERROR CREATING OPENING: %v", err)
		defaultMessages.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	defaultMessages.SendSuccess(context, "Opening created successfully", true)
}
