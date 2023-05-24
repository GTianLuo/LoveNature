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

			userGroup.POST("/logout/:nickName", middleware.CheckLoginStatus(), v1.Logout)

			userGroup.PATCH("/password/:nickName", middleware.CheckLoginStatus(), v1.UpdatePassword)

			userGroup.PATCH("/edit/nickName/:nickName", middleware.CheckLoginStatus(), v1.UpdateNickName)
			//userGroup.GET("/me")
			userGroup.POST("/icon/:nickName", middleware.CheckLoginStatus(), v1.UploadIcon)
		}

		userInfoGroup := v1Group.Group("/userInfo", middleware.CheckLoginStatus())
		{
			userInfoGroup.GET("/me/:nickName", v1.GetMeInfo)
			userInfoGroup.PATCH("/edit/sex/:nickName", v1.UpdateSex)
			userInfoGroup.PATCH("/edit/address/:nickName", v1.UpdateAddress)
			userInfoGroup.PATCH("/edit/introduction/:nickName", v1.UpdateIntroduction)
		}

		petGroup := v1Group.Group("/pet")
		{
			petGroup.POST("admin/petInfo", v1.PostPetInfo)
			petGroup.POST("admin/petInfoPic", v1.PostPetInfoPic)
			petGroup.GET("keywordList/:keyword", v1.SearchByKeyword)
			petGroup.GET("petInfo/:name", v1.GetPetInfo)
		}

		blogGroup := v1Group.Group("/blog", middleware.RefreshToken())
		{
			blogGroup.POST("/content/:nickName", middleware.CheckLoginStatus(), v1.PostBlog)
			blogGroup.GET("/content/:keyword/:page", v1.SearchBlog)
			blogGroup.GET("/content/list/:way/:page", v1.GetBlogsList)
		}
	}
	return r
}
