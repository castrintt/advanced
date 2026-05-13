package config

import (
	"goportunities/schemas"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	databasePath = "./db/main.db"
)

func checkIfDatabaseFileExists() (bool, error) {
	_, err := os.Stat(databasePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func ensureDatabaseFileOnDisk() error {
	exists, err := checkIfDatabaseFileExists()
	if exists && err == nil {
		return nil
	}
	logger.Info("Creating database file...")
	if err = os.MkdirAll(filepath.Dir(databasePath), os.ModePerm); err != nil {
		logger.Errorf("ERROR CREATING DATABASE FILE: %v", err)
		return err
	}
	f, err := os.Create(databasePath)
	if err != nil {
		logger.Errorf("ERROR CREATING DATABASE FILE: %v", err)
		return err
	}
	return f.Close()
}

func InitializeSqlite() (*gorm.DB, error) {
	logger = GetLogger("SQLITE")

	if err := ensureDatabaseFileOnDisk(); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

	if err != nil {
		logger.Errorf("ERROR OPENING SQLITE DATABASE: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("ERROR MIGRATING SCHEMAS: %v", err)
		return nil, err
	}

	return db, nil
}
