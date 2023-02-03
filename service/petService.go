package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"lovenature/dao"
	"lovenature/dto"
	"lovenature/model"
	"lovenature/pkg/e"
	"lovenature/pkg/util"
	"mime/multipart"
)

type PetService struct {
	Name         string `form:"name"`
	Introduction string `form:"introduction"`
	KeyWord      string `form:"keyword"`
}

func NewPetService() *PetService {
	return &PetService{}
}

func (s *PetService) PostPetInfo(ctx context.Context) *dto.Result {
	petDao := dao.NewPetDao(ctx)
	pet := &model.Pet{
		Name:         s.Name,
		Introduction: s.Introduction,
		KeyWord:      s.KeyWord,
	}
	if err := petDao.CreatePetInfo(pet); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, nil)
}

func (s *PetService) PostPetInfoPic(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) *dto.Result {
	petDao := dao.NewPetDao(ctx)
	//校验文件
	if header.Size > (8 << 18) {
		return dto.Fail(e.IconTooBig, nil)
	}
	if typ := header.Header.Get("Content-Type"); typ != "image/png" &&
		typ != "image/gif" &&
		typ != "image/jpeg" &&
		typ != "image/jpg" &&
		typ != "application/octet-stream" &&
		typ != "image/bmp" {
		return dto.Fail(e.WrongPictureFormat, nil)
	}
	//上传图片
	url, err := util.UploadImg(file, header.Size)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	//修改数据库
	if err := petDao.UploadPic(s.Name, url); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, nil)
}

// SearchByKeyWord 通过关键字搜索宠物列表
func (s *PetService) SearchByKeyWord(ctx *gin.Context, keyword string) *dto.Result {
	petDao := dao.NewPetDao(ctx)
	pets, err := petDao.SearchByKeyword(keyword)
	if err != nil {
		return dto.Fail(e.Error, nil)
	}
	petDaos := dto.BuildPetDaoList(pets)
	if len(petDaos) == 0 {
		return dto.Fail(e.NotFoundInformation, nil)
	}
	return dto.Success(e.Success, petDaos)
}
