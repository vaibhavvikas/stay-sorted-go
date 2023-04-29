package house_service

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vaibhavvikas/stay-sorted/constants"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
	"github.com/vaibhavvikas/stay-sorted/repository/house_repository"
	"github.com/vaibhavvikas/stay-sorted/utils"
	"gorm.io/gorm"
)

type HouseServiceImpl struct {
	houseRepository house_repository.HouseRepositoryImpl
}

func (h HouseServiceImpl) CreateHouse(ctx *gin.Context, houseReq models.CreateHouseReq) (entities.House, error) {
	var house entities.House = entities.House{
		UserPid:       ctx.MustGet("user_pid").(string),
		Name:          houseReq.HouseName,
		Description:   houseReq.Description,
		Address:       houseReq.Address,
		City:          houseReq.City,
		State:         houseReq.State,
		Country:       houseReq.Country,
		Pincode:       houseReq.Pincode,
		NumBedrooms:   houseReq.Bedrooms,
		NumBathrooms:  houseReq.Bathrooms,
		Price:         houseReq.Price,
		SquareFootage: houseReq.SquareFootage,
		Amenities:     strings.Join(houseReq.Amenities, ", "),
	}

	house, err := h.houseRepository.CreateHouse(ctx, house)
	if err != nil {
		return house, errors.Wrap(err, "[CreateHouse]")
	}

	var housePictures []entities.HousePicture
	for i := 0; i < len(houseReq.Pictures); i++ {
		housePicture := entities.HousePicture{
			Pid:        utils.GeneratePid(constants.Prefixes.HousePicture, 20),
			HousePid:   house.Pid,
			PictureURL: houseReq.Pictures[i],
		}
		housePictures = append(housePictures, housePicture)
	}

	err = h.houseRepository.CreatePictures(ctx, housePictures)
	if err != nil {
		return house, errors.Wrap(err, "[CreateHouse]")
	}

	house.HousePictures = housePictures
	return house, errors.Wrap(err, "[CreateHouse]")
}

func (h HouseServiceImpl) GetHouseByPid(ctx *gin.Context, housePid string) (entities.House, error) {
	house, err := h.houseRepository.GetHouseByPid(ctx, housePid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return house, getInvalidHouseError(housePid)
		}
		return house, errors.Wrap(err, "[GetHouseByPid]")
	}
	return house, errors.Wrap(err, "[GetHouseByPid]")
}

func getInvalidHouseError(pid string) error {
	return &models.AppError{
		Message: "The house with following pid does not exist in the system: " + pid,
		Param:   "house_pid",
	}
}

func (h HouseServiceImpl) GetAllHouses(ctx *gin.Context) ([]entities.House, error) {
	houses, err := h.houseRepository.GetAllHouses(ctx)
	if err != nil {
		return houses, errors.Wrap(err, "[GetAllHouses]")
	}
	return houses, errors.Wrap(err, "[GetAllHouses]")
}

// func (h HouseServiceImpl) UpdateHouse(ctx *gin.Context, housePid string) (entities.House, error) {}
