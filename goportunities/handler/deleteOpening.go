package handler

import (
	"goportunities/handler/defaultMessages"
	"goportunities/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(context *gin.Context) {
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

	if err := db.Delete(&schemas.Opening{}, id).Error; err != nil {
		defaultMessages.SendError(context, http.StatusInternalServerError, "Failed to delete opening")
		logger.Errorf("ERROR DELETING OPENING: %v", err)
		return
	}
	defaultMessages.SendSuccess(context, "Opening deleted successfully", true)
}
