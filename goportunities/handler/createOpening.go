package handler

import (
	Requests "goportunities/handler/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := Requests.CreateOpeningRequest{}
	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("ERROR CREATING OPENING: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create opening"})
		return
	}
}
