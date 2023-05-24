package dto

import "lovenature/model"

type UserDto struct {
	Email    string `json:"email,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Icon     string `json:"icon"`
	Token    string `json:"token,omitempty"`
	//HasLikedBlog []string `json:"hasLikedBlog"`
}

func BuildUser(user *model.User, token string) *UserDto {
	return &UserDto{
		Email:    user.Email,
		NickName: user.NickName,
		Token:    token,
		Icon:     user.Icon,
	}
}

func BuildUserList(user []model.User) []*UserDto {
	var users []*UserDto
	for _, user := range user {
		userDto := &UserDto{
			Email:    user.Email,
			NickName: user.NickName,
			Icon:     user.Icon,
		}
		users = append(users, userDto)
	}
	return users
}
