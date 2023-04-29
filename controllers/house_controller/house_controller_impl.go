package house_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/services/house_service"
	"github.com/vaibhavvikas/stay-sorted/transformers"
	"github.com/vaibhavvikas/stay-sorted/validators/house_validator"
)

type HouseControllerImpl struct {
	houseService house_service.HouseServiceImpl
}

func (h HouseControllerImpl) CreateHouse(ctx *gin.Context) {
	houseReq, err := house_validator.ValidateCreateHouseRequest(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	house, err := h.houseService.CreateHouse(ctx, houseReq)
	if err != nil {
		ctx.Error(err)
		return
	}

	houseResponse := transformers.GetHouseResponse(house)
	ctx.IndentedJSON(http.StatusCreated, houseResponse)
}

func (h HouseControllerImpl) GetHouse(ctx *gin.Context) {
	housePid, err := house_validator.ValidateGetHouseRequest(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	house, err := h.houseService.GetHouseByPid(ctx, housePid)
	if err != nil {
		ctx.Error(err)
		return
	}

	houseResponse := transformers.GetHouseResponse(house)
	ctx.IndentedJSON(http.StatusCreated, houseResponse)
}

func (h HouseControllerImpl) GetAllHouses(ctx *gin.Context) {
	houses, err := h.houseService.GetAllHouses(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	housesResponse := transformers.GetAllHousesResponse(houses)
	ctx.IndentedJSON(http.StatusCreated, housesResponse)
}
