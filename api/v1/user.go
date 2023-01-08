package v1

import (
	"github.com/gin-gonic/gin"
	"lovenature/dto"
	"lovenature/service"
	"net/http"
)

func SendCode(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.SendCode())
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func Register(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.Register(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func LoginByPassword(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.LoginByPassword(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func LoginByCode(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.LoginByCode(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UpdatePassword(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.UpdatePassword(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func Logout(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.Logout(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func UploadIcon(ctx *gin.Context) {
	userService := service.NewUserService()
	if err := ctx.ShouldBind(userService); err == nil {
		ctx.JSON(http.StatusOK, userService.UploadIcon(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}
