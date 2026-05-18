package handler

import (
	"goportunities/handler/defaultMessages"
	requests "goportunities/handler/request"
	"goportunities/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	paramId := c.Param("id")
	if paramId == "" {
		defaultMessages.SendError(c, http.StatusBadRequest, "ID is required")
		return
	}

	id64, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		defaultMessages.SendError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	id := uint(id64)

	request := requests.UpdateOpeningRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		defaultMessages.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	entity := request.ToEntity()

	if err := db.Model(&schemas.Opening{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		defaultMessages.SendError(c, http.StatusInternalServerError, "Failed to update opening")
		return
	}

	defaultMessages.SendSuccess(c, "Opening updated successfully", true)
}
