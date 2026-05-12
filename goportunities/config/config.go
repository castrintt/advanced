package config

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

func Initialize() error {
	return nil
}
