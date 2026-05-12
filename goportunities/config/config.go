package config

import (
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	logger *Logger
)

func Initialize() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
