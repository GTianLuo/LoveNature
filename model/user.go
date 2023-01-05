package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(25) not null"`
	Email    string `gorm:"type:varchar(20) not null unique"`
	Password string `gorm:"type:varchar(255)"`
	Sex      int    `gorm:"type:int default 2"`
	Icon     string `gorm:"type:varchar(256) default 'http://rnyrwpase.bkt.clouddn.com/default.jpg'"`
}

func (*User) TableName() string {
	return "user"
}