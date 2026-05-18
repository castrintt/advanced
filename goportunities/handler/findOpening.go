package handler

import (
	"goportunities/handler/defaultMessages"
	"goportunities/handler/response"
	"goportunities/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindOpeningHandler(context *gin.Context) {
	paramId := context.Param("id")

	if paramId == "" {
		defaultMessages.SendError(context, http.StatusBadRequest, "ID is required")
		return
	}

	id64, err := strconv.ParseUint(paramId, 10, 64)

	if err != nil {
		defaultMessages.SendError(context, http.StatusBadRequest, "Invalid ID")
		return
	}

	id := uint(id64)

	var entity schemas.Opening
	if err := db.First(&entity, id).Error; err != nil {
		defaultMessages.SendError(context, http.StatusNotFound, "Opening not found")
		logger.Errorf("ERROR FINDING OPENING: %v", err)
		return
	}

	var out response.OpeningResponse
	wrapped := out.FromEntity(&entity)
	defaultMessages.SendSuccess(context, "Opening found successfully", wrapped)
}
