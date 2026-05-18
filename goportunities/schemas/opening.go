package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `gorm:"not null"`
	Company  string `gorm:"not null"`
	Location string `gorm:"not null"`
	Remote   bool   `gorm:"not null"`
	Link     string `gorm:"not null"`
	Salary   int64  `gorm:"not null"`
}
