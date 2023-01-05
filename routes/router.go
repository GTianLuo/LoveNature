package routes

import (
	"github.com/gin-gonic/gin"
	v1 "lovenature/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//v1
	v1Group := r.Group("api/v1")
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/code", v1.SendCode)
			userGroup.POST("/register", v1.Register)
			userGroup.POST("/login/code", v1.LoginByCode)
			userGroup.POST("/login/password", v1.LoginByPassword)
		}
	}
	return r
}
