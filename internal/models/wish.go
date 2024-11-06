package models

import (
	"gorm.io/gorm"
)

type Wish struct {
	gorm.Model
	UserID uint

	Text string `gorm:"type:text;not null"`
}
