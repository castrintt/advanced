package handler

import (
	Requests "goportunities/handler/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := Requests.CreateOpeningRequest{}
	context.BindJSON(&request)

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("ERROR CREATING OPENING: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create opening"})
		return
	}
}
