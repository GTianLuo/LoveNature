package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"lovenature/conf"
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
	petDaos := *dto.BuildPetDtoList(pets)
	if len(petDaos) == 0 {
		return dto.Fail(e.NotFoundInformation, nil)
	}
	return dto.Success(e.Success, petDaos)
}

func (s *PetService) GetPetInfo(ctx *gin.Context, name string) *dto.Result {
	petDao := dao.NewPetDao(ctx)
	redisClient := conf.NewRedisClient()
	//先从redis中查找
	if redisClient.Exists(e.PetHotData+name).Val() > 0 {
		//刷新过期时间
		redisClient.Expire(e.PetHotData+name, e.PetHotDataDDL)
		if petDtoMap, err := redisClient.HGetAll(e.PetHotData + name).Result(); err != nil {
			return dto.Fail(e.Error, err)
		} else {
			var petDto dto.PetDto
			util.MapToStruct(petDtoMap, &petDto)
			return dto.Success(e.Success, petDto)
		}
	}
	//redis中查到数据，刷新过期时间
	//未查到数据，则查询mysql
	pet, err := petDao.GetPetInfoByName(name)
	if err != nil {
		return dto.Fail(e.Error, err)
	} else if pet.Name == "" {
		return dto.Fail(e.NotFoundInformation, nil)
	}
	petDto := dto.BuildPetDto(pet)
	//将数据保存到redis
	redisClient.HMSet(e.PetHotData+name, util.StructToMap(petDto))
	redisClient.Expire(e.PetHotData+name, e.PetHotDataDDL)
	return dto.Success(e.Success, petDto)
}
