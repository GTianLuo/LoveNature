package v1

import (
	"github.com/gin-gonic/gin"
	"lovenature/dto"
	"lovenature/pkg/e"
	"lovenature/service"
	"net/http"
)

func PostPetInfo(ctx *gin.Context) {
	petService := service.NewPetService()
	if err := ctx.ShouldBind(petService); err == nil {
		ctx.JSON(http.StatusOK, petService.PostPetInfo(ctx))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, err))
	}
}

func PostPetInfoPic(ctx *gin.Context) {
	petService := service.NewPetService()
	file, header, err := ctx.Request.FormFile("picture")
	if err == nil {
		petService.Name = ctx.PostForm("name")
		ctx.JSON(http.StatusOK, petService.PostPetInfoPic(ctx, file, header))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(e.InvalidParam, err))
	}
}

func SearchByKeyword(ctx *gin.Context) {
	keyword := ctx.Param("keyword")
	petService := service.NewPetService()
	if keyword != "" {
		ctx.JSON(http.StatusOK, petService.SearchByKeyWord(ctx, keyword))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(e.InvalidParam, nil))
	}
}

func GetPetInfo(ctx *gin.Context) {
	petService := service.NewPetService()
	name := ctx.Param("name")
	if name != "" {
		ctx.JSON(http.StatusOK, petService.GetPetInfo(ctx, name))
	} else {
		ctx.JSON(http.StatusBadRequest, dto.Fail(e.InvalidParam, nil))
	}
}
