package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID        uint `gorm:"not null"`
	City          string
	Introduction  string
	Followee      uint
	Interested    string
	Fans          uint
	NotesNumber   uint
	CollectNumber uint
}

func (*UserInfo) TableName() string {
	return "user_info"
}
