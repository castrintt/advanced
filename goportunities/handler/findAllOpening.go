package handler

import (
	"goportunities/handler/defaultMessages"
	"goportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllOpeningHandler(c *gin.Context) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		defaultMessages.SendError(c, http.StatusInternalServerError, "Failed to find all openings")
		return
	}

	defaultMessages.SendSuccess(c, "Openings found successfully", openings)
}
