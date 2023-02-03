package dto

import "lovenature/model"

type PetDto struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Picture      string `json:"picture,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

func BuildPetDtoList(pets []model.Pet) *[]PetDto {
	petDaoList := make([]PetDto, len(pets))
	for i, pet := range pets {
		petDao := PetDto{
			Name:         pet.Name,
			Image:        pet.Image,
			Picture:      pet.Picture,
			Introduction: pet.Introduction,
		}
		petDaoList[i] = petDao
	}
	return &petDaoList
}

func BuildPetDto(pet *model.Pet) *PetDto {
	return &PetDto{
		Name:         pet.Name,
		Image:        pet.Image,
		Picture:      pet.Picture,
		Introduction: pet.Introduction,
	}
}
