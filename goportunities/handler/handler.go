package handler

import (
	"goportunities/config"
	"goportunities/handler/requests"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	requests.InitializeRequests()
	logger = config.GetLogger("HANDLER")
	logger.Info("Initializing handler...")
	db = config.GetSqlite()
}
