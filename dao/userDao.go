package dao

import (
	"context"
	"gorm.io/gorm"
	"lovenature/conf"
	"lovenature/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		db: conf.NewDBClient(ctx),
	}
}

func (dao *UserDao) IsExistByEmail(email string) bool {
	var count int64
	if dao.db.Model(&model.User{}).Where("email = ?", email).Count(&count); count == 1 {
		return true
	}
	return false
}

func (dao *UserDao) CreateUser(user *model.User) error {

	return dao.db.Transaction(func(tx *gorm.DB) error {
		//插入user
		if err := dao.db.Create(user).Error; err != nil {
			return err
		}
		userInfo := &model.UserInfo{
			UserID: user.ID,
		}
		return dao.db.Create(&userInfo).Error
	})
}

func (dao *UserDao) GetUser(email string) *model.User {
	u := &model.User{}
	dao.db.Where("email = ?", email).Find(u)
	return u
}

func (dao *UserDao) UpdatePassword(nickName string, password string) error {
	return dao.db.Model(&model.User{}).Where("nick_name = ?", nickName).Update("password", password).Error
}

func (dao *UserDao) UploadIcon(nickName string, url string) error {
	return dao.db.Model(&model.User{}).Where("nick_name = ?", nickName).Update("icon", url).Error
}
