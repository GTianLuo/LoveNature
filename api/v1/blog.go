package v1

import (
	"github.com/gin-gonic/gin"
	"lovenature/dto"
	"lovenature/pkg/e"
	"lovenature/service"
	"net/http"
	"strconv"
)

func PostBlog(ctx *gin.Context) {
	blogService := service.NewBlogService()
	blogService.Email = ctx.Keys["user"].(*dto.UserDto).Email
	if err := ctx.ShouldBind(blogService); err == nil {
		ctx.JSON(http.StatusOK, blogService.PostBlog(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func SearchBlog(ctx *gin.Context) {
	keyword := ctx.Param("keyword")
	page, err := strconv.Atoi(ctx.Param("page"))
	blogService := service.NewBlogService()
	if keyword != "" && err == nil && page > 0 {
		ctx.JSON(http.StatusOK, blogService.SearchByKeyWord(ctx, keyword, page))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(e.InvalidParam, nil))
	}
}

func GetBlogsList(ctx *gin.Context) {
	way := ctx.Param("way")
	page, err := strconv.Atoi(ctx.Param("page"))
	blogService := service.NewBlogService()
	if way != "" && err == nil && page > 0 {
		ctx.JSON(http.StatusOK, blogService.GetBlogList(ctx, way, page))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(e.InvalidParam, err))
	}
}
