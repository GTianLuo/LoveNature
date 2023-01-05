package dto

import "lovenature/model"

type UserDto struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Sex      int    `json:"sex"`
	Icon     string `json:"icon"`
	Token    string `json:"token,omitempty"`
}

func BuildUser(user *model.User, token string) *UserDto {
	return &UserDto{
		Email:    user.Email,
		NickName: user.NickName,
		Sex:      user.Sex,
		Token:    token,
		Icon:     user.Icon,
	}
}
