package service

import (
	"github.com/gin-gonic/gin"
	"lovenature/dao"
	"lovenature/dto"
	"lovenature/pkg/e"
)

type UserInfoService struct {
	NickName     string `form:"nickName"`
	Token        string `form:"token"`
	Address      string `form:"address"`
	Sex          int    `form:"sex"`
	Introduction string `form:"introduction"`
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (s *UserInfoService) UpdateAddress(ctx *gin.Context) *dto.Result {

	userInfoDao := dao.NewUserInfoDao(ctx)
	if s.Address == "" {
		return dto.Fail(e.NilAddress, nil)
	}
	if err := userInfoDao.UpdateCity(s.NickName, s.Address); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "修改成功")
}

func (s *UserInfoService) UpdateSex(ctx *gin.Context) *dto.Result {
	userInfoDao := dao.NewUserInfoDao(ctx)
	if s.Sex != 1 && s.Sex != 2 && s.Sex != 0 {
		return dto.Fail(e.InvalidSex, nil)
	}
	if err := userInfoDao.UpdateSex(s.NickName, s.Sex); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "修改成功")
}

func (s *UserInfoService) UpdateIntroduction(ctx *gin.Context) *dto.Result {
	userInfoDao := dao.NewUserInfoDao(ctx)
	if len(s.Introduction) > 150 {
		return dto.Fail(e.IntroductionIsTooLong, nil)
	}
	if err := userInfoDao.UpdateIntroduction(s.NickName, s.Introduction); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "修改成功")
}

func (s *UserInfoService) GetMeInfo(ctx *gin.Context, nickName, token string) *dto.Result {
	userInfoDao := dao.NewUserInfoDao(ctx)
	userInfo, err := userInfoDao.GetMeInfo(nickName)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	user, _ := ctx.Get("user")
	return dto.Success(e.Success, dto.BuildUserInfo(user.(*dto.UserDto), userInfo))
}
