package dto

import "lovenature/model"

type PetDao struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Picture      string `json:"picture,omitempty"`
	Introduction string `json:"introduction,omitempty"`
}

func BuildPetDaoList(pets []model.Pet) []PetDao {
	petDaoList := make([]PetDao, len(pets))
	for i, pet := range pets {
		petDao := PetDao{
			Name:         pet.Name,
			Image:        pet.Image,
			Picture:      pet.Picture,
			Introduction: pet.Introduction,
		}
		petDaoList[i] = petDao
	}
	return petDaoList
}
