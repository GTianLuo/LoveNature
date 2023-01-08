package middleware

import (
	"github.com/gin-gonic/gin"
	"lovenature/conf"
	"lovenature/dto"
	"lovenature/pkg/e"
	"lovenature/pkg/util"
	"net/http"
)

func CheckLoginStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redisClient := conf.NewRedisClient()
		//获取token和nickName
		token := ctx.PostForm("token")
		nickName := ctx.PostForm("nickName")
		if token == "" || nickName == "" {
			//用户未登录
			ctx.Abort()
			ctx.JSON(http.StatusOK, dto.Fail(e.UserNotLogin, nil))
			return
		}
		//获取用户信息
		result, err := redisClient.HGetAll(e.UserLoginInfo + nickName).Result()
		if err != nil || result["Token"] != token {
			//用户未登录
			ctx.Abort()
			ctx.JSON(http.StatusOK, dto.Fail(e.UserNotLogin, err))
			return
		}
		user := &dto.UserDto{}
		util.MapToStruct(result, user)
		ctx.Set("user", user)
	}
}
