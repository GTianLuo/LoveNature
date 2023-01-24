package middleware

import (
	"github.com/gin-gonic/gin"
	"lovenature/conf"
	"lovenature/pkg/e"
)

func RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		nickName := ctx.Query("nickName")
		if token == "" || nickName == "" {
			return
		}
		redisClient := conf.NewRedisClient()
		//用户未登录
		if cnt := redisClient.Exists(e.UserLoginInfo + nickName).Val(); cnt <= 1 {
			return
		}
		//获取用户信息并比对token
		if user, err := redisClient.HGetAll(e.UserLoginInfo + nickName).Result(); err != nil || user["Token"] != token {
			return
		}

		redisClient.Expire(e.UserLoginInfo+nickName, e.UserLoginInfoTTL)
	}
}
