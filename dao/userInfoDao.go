package dao

import (
	"context"
	"gorm.io/gorm"
	"lovenature/conf"
	"lovenature/model"
)

type UserInfoDao struct {
	db *gorm.DB
}

func NewUserInfoDao(ctx context.Context) *UserInfoDao {
	return &UserInfoDao{
		db: conf.NewDBClient(ctx),
	}
}

func (dao *UserInfoDao) UpdateCity(name string, address string) error {
	return dao.db.Exec("UPDATE user_info a JOIN user b ON a.user_id = b.id  SET a.address = ? WHERE b.nick_name = ?", address, name).Error
}

func (dao *UserInfoDao) UpdateSex(name string, sex int) error {
	return dao.db.Exec("UPDATE user_info a JOIN user b ON a.user_id = b.id  SET a.sex = ? WHERE b.nick_name = ?", sex, name).Error
}

func (dao *UserInfoDao) UpdateIntroduction(name string, introduction string) error {
	return dao.db.Exec("UPDATE user_info a JOIN user b ON a.user_id = b.id  SET a.introduction = ? WHERE b.nick_name = ?", introduction, name).Error
}

func (dao *UserInfoDao) GetMeInfo(name string) (*model.UserInfo, error) {
	userInfo := &model.UserInfo{}
	//err := dao.db.Where("nick_name = ?", name).Find(userInfo).Error
	//"select * from user_info a join user b on a.user_id = b.id where a.nick_name = ?"
	err := dao.db.
		Table("user_info").
		Joins("join user on user_info.user_id = user.id").
		Where("user.nick_name = ?", name).
		Find(userInfo).Error
	return userInfo, err
}
