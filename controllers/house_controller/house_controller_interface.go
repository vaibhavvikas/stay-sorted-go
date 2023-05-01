package house_controller

import "github.com/gin-gonic/gin"

type HouseControllerInterface interface {
	CreateHouse(ctx *gin.Context)
	GetHouse(ctx *gin.Context)
	GetAllHouses(ctx *gin.Context)
	GetFilteredHouses(ctx *gin.Context)
}
