package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitializeAppConfig() error {
	var err error
	db, err = InitializeSqlite()
	if err != nil {
		return err
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
