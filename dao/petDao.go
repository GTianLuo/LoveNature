package dao

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"lovenature/conf"
	"lovenature/model"
)

type PetDao struct {
	db *gorm.DB
}

func NewPetDao(ctx context.Context) *PetDao {
	return &PetDao{
		db: conf.NewDBClient(ctx),
	}
}

func (dao *PetDao) CreatePetInfo(pet *model.Pet) error {
	return dao.db.Create(pet).Error
}

func (dao *PetDao) UploadPic(name string, url string) error {
	//先读取现在数据库有的图片
	pet := &model.Pet{}
	if err := dao.db.Select("p_picture").Where("p_name = ?", name).Find(pet).Error; err != nil {
		return err
	}
	//将json转换为数组，再将这张图片上传
	picJson := []byte(pet.Picture)
	var picArr []string
	if err := json.Unmarshal(picJson, &picArr); err != nil && pet.Picture != "" {
		return err
	}
	picArr = append(picArr, url)
	//将数组转换为json更新原先的pic
	picJson, err := json.Marshal(picArr)
	if err != nil {
		return err
	}
	return dao.db.Model(&model.Pet{}).Where("p_name = ?", name).Update("p_picture", string(picJson)).Error

}

func (dao *PetDao) SearchByKeyword(keyword string) ([]model.Pet, error) {
	var pets []model.Pet
	err := dao.db.Select("p_name", "p_img").Where("POSITION(? IN `keyword`)", keyword).Find(&pets).Error
	return pets, err
}
