package house_service

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
)

type HouseServiceInterface interface {
	CreateHouse(ctx *gin.Context, houseReq models.CreateHouseReq) (entities.House, error)
	GetHouseByPid(ctx *gin.Context, housePid string) (entities.House, error)
	GetAllHouses(ctx *gin.Context) ([]entities.House, error)
	// UpdateHouse(ctx *gin.Context, housePid string) (entities.House, error)
	// Remove(ctx *gin.Context, housePid string) (bool, string)
}
