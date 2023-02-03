package model

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name         string `gorm:"column:p_name;unique"`
	Image        string `gorm:"column:p_img"`
	Picture      string `gorm:"type:text;column:p_picture"`
	Introduction string `gorm:"column:p_intro"`
	KeyWord      string `gorm:"column:keyword;index"`
}

func (*Pet) TableName() string {
	return "pet"
}
