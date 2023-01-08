package routes

import (
	"github.com/gin-gonic/gin"
	v1 "lovenature/api/v1"
	"lovenature/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//v1
	v1Group := r.Group("api/v1", middleware.RefreshToken())
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.POST("/code", v1.SendCode)
			userGroup.POST("/register", v1.Register)
			userGroup.POST("/login/code", v1.LoginByCode)
			userGroup.POST("/login/password", v1.LoginByPassword)

			userGroup.POST("/logout", middleware.CheckLoginStatus(), v1.Logout)
			userGroup.PATCH("/password", middleware.CheckLoginStatus(), v1.UpdatePassword)
			//userGroup.GET("/me")
			userGroup.POST("/icon", middleware.CheckLoginStatus(), v1.UploadIcon)
		}
	}
	return r
}
