package dto

import "lovenature/model"

type UserInfoDto struct {
	Email         string `json:"email"`
	NickName      string `json:"nickName"`
	Icon          string `json:"icon"`
	Address       string `json:"address"`
	Sex           int    `json:"sex"`
	Introduction  string `json:"introduction"`
	Followee      uint   `json:"followee"`
	Fans          uint   `json:"fans"`
	NotesNumber   uint   `json:"notesNumber"`
	CollectNumber uint   `json:"notes"`
}

func BuildUserInfo(user *UserDto, userInfo *model.UserInfo) *UserInfoDto {
	return &UserInfoDto{
		Email:         user.Email,
		NickName:      user.NickName,
		Icon:          user.Icon,
		Address:       userInfo.Address,
		Sex:           userInfo.Sex,
		Introduction:  userInfo.Introduction,
		Followee:      userInfo.Followee,
		Fans:          userInfo.Fans,
		NotesNumber:   userInfo.NotesNumber,
		CollectNumber: userInfo.CollectNumber,
	}
}
