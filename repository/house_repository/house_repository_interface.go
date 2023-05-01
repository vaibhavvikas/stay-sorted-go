package house_repository

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
)

type HouseRepository interface {
	CreateHouse(ctx *gin.Context, house entities.House) (entities.House, error)
	UpdateHouse(ctx *gin.Context, house entities.House) (entities.House, error)
	GetHouseByPid(ctx *gin.Context, pid string) (entities.House, error)
	GetFilteredHouses(ctx *gin.Context, filters models.HouseFilter) ([]entities.House, error)
	RemoveHouse(ctx *gin.Context, pid string) error
}
