package v1

import (
	"github.com/gin-gonic/gin"
	"lovenature/dto"
	"lovenature/service"
	"net/http"
)

func UpdateAddress(ctx *gin.Context) {
	userInfoService := service.NewUserInfoService()
	if err := ctx.ShouldBind(userInfoService); err == nil {
		userInfoService.NickName = ctx.Param("nickName")
		ctx.JSON(http.StatusOK, userInfoService.UpdateAddress(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UpdateSex(ctx *gin.Context) {
	userInfoService := service.NewUserInfoService()
	if err := ctx.ShouldBind(userInfoService); err == nil {
		userInfoService.NickName = ctx.Param("nickName")
		ctx.JSON(http.StatusOK, userInfoService.UpdateSex(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UpdateIntroduction(ctx *gin.Context) {
	userInfoService := service.NewUserInfoService()
	if err := ctx.ShouldBind(userInfoService); err == nil {
		userInfoService.NickName = ctx.Param("nickName")
		ctx.JSON(http.StatusOK, userInfoService.UpdateIntroduction(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func GetMeInfo(ctx *gin.Context) {
	userInfoService := service.NewUserInfoService()
	nickName := ctx.Param("nickName")
	token := ctx.GetHeader("token")
	if token != "" && nickName != "" {
		ctx.JSON(http.StatusOK, userInfoService.GetMeInfo(ctx, nickName, token))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, nil))
	}
}
