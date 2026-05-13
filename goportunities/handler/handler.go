package handler

import (
	"goportunities/config"
	Requests "goportunities/handler/request"
	Responses "goportunities/handler/response"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	Requests.InitializeRequests()
	Responses.InitializeResponse()
	logger = config.GetLogger("HANDLER")
	db = config.GetSqlite()
}
