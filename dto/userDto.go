package dto

import "lovenature/model"

type UserDto struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
	Token    string `json:"token,omitempty"`
}

func BuildUser(user *model.User, token string) *UserDto {
	return &UserDto{
		Email:    user.Email,
		NickName: user.NickName,
		Token:    token,
		Icon:     user.Icon,
	}
}
