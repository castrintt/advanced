package handler

import (
	"goportunities/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("HANDLER")
	logger.Info("Initializing handler...")
	db = config.GetSqlite()
}
