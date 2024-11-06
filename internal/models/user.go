package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	TelegramId int64  `gorm:"unique;not null"`

	Wishes []Wish
}
