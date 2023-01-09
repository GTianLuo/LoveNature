package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID        uint `gorm:"not null"`
	Address       string
	Sex           int `gorm:"default:2"`
	Introduction  string
	Followee      uint
	Fans          uint
	Interested    string
	NotesNumber   uint
	CollectNumber uint
}

func (*UserInfo) TableName() string {
	return "user_info"
}
