package transformers

import (
	"strings"

	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
)

func GetHouseResponse(house entities.House) models.HouseResponse {
	var pictures []string
	for i := 0; i < len(house.HousePictures); i++ {
		pictures = append(pictures, house.HousePictures[i].PictureURL)
	}

	return models.HouseResponse{
		HousePid:      house.Pid,
		HouseName:     house.Name,
		Address:       house.Address,
		City:          house.City,
		State:         house.State,
		Country:       house.Country,
		Pincode:       house.Pincode,
		Description:   house.Description,
		Bedrooms:      house.NumBedrooms,
		Bathrooms:     house.NumBathrooms,
		Price:         house.Price,
		SquareFootage: house.SquareFootage,
		Amenities:     strings.Split(house.Amenities, ", "),
		Pictures:      pictures,
	}
}

func GetAllHousesResponse(houses []entities.House) []models.HouseResponse {
	var housesResp []models.HouseResponse
	for i := 0; i < len(houses); i++ {
		housesResp = append(housesResp, GetHouseResponse(houses[i]))
	}
	return housesResp
}
