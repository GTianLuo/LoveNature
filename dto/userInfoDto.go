package dto

import "lovenature/model"

type UserInfoDto struct {
	Email          string `json:"email"`
	NickName       string `json:"nickName"`
	Icon           string `json:"icon"`
	Address        string `json:"address"`
	Introduction   string `json:"introduction""`
	Sex            int    `json:"sex"`
	Followee       uint   `json:"followee"`
	Fans           uint   `json:"fans"`
	NotesNumber    uint   `json:"notesNumber"`
	CollectNumber  uint   `json:"collectNumber"`
	GetLikesNumber int64  `json:"getLikesNumber"`
}

func BuildUserInfo(user *UserDto, userInfo *model.UserInfo) *UserInfoDto {
	return &UserInfoDto{
		Email:          user.Email,
		NickName:       user.NickName,
		Icon:           user.Icon,
		Address:        userInfo.Address,
		Introduction:   userInfo.Introduction,
		Sex:            userInfo.Sex,
		Followee:       userInfo.Followee,
		Fans:           userInfo.Fans,
		NotesNumber:    userInfo.NotesNumber,
		CollectNumber:  userInfo.CollectNumber,
		GetLikesNumber: userInfo.GetLikesNumber,
	}
}
