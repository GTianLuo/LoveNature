package dao

import (
	"context"
	"gorm.io/gorm"
	"lovenature/conf"
	"lovenature/model"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		DB: conf.NewDBClient(ctx),
	}
}

func (dao *UserDao) IsExistByEmail(email string) bool {
	var count int64
	if dao.DB.Model(&model.User{}).Where("email = ?", email).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(user).Error
}

func (dao *UserDao) GetUser(email string) *model.User {
	u := &model.User{}
	dao.DB.Where("email = ?", email).Find(u)
	return u
}
