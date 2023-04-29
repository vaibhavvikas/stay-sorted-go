package house_repository

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vaibhavvikas/stay-sorted/constants"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
	"github.com/vaibhavvikas/stay-sorted/utils"
	"gorm.io/gorm"
)

type HouseRepositoryImpl struct{}

func (h HouseRepositoryImpl) CreateHouse(ctx *gin.Context, house entities.House) (entities.House, error) {
	db := ctx.MustGet("DB").(*gorm.DB)
	house.Pid = utils.GeneratePid(constants.Prefixes.House, 25)

	err := db.Create(&house).Error
	if err != nil {
		return house, errors.Wrap(err, "[HouseRepositoryImpl][CreateHouse]")
	}
	return house, errors.Wrap(err, "[HouseRepositoryImpl][CreateHouse]")
}

func (h HouseRepositoryImpl) UpdateHouse(ctx *gin.Context, house entities.House) (entities.House, error) {
	db := ctx.MustGet("DB").(*gorm.DB)
	err := db.Save(&house).Error
	if err != nil {
		return house, errors.Wrap(err, "[HouseRepositoryImpl][UpdateHouse]")
	}
	return house, errors.Wrap(err, "[HouseRepositoryImpl][UpdateHouse]")
}

func (h HouseRepositoryImpl) GetHouseByPid(ctx *gin.Context, pid string) (entities.House, error) {
	var house entities.House
	db := ctx.MustGet("DB").(*gorm.DB)

	err := db.Preload("HousePictures").Where("house_pid = ?", pid).Take(&house).Error
	if err != nil {
		return house, errors.Wrap(err, "[HouseRepositoryImpl][GetHouseByPid]")
	}
	return house, errors.Wrap(err, "[HouseRepositoryImpl][GetHouseByPid]")
}

func (h HouseRepositoryImpl) GetAllHouses(ctx *gin.Context) ([]entities.House, error) {
	var houses []entities.House
	db := ctx.MustGet("DB").(*gorm.DB)

	err := db.Preload("HousePictures").Find(&houses).Error
	if err != nil {
		return houses, errors.Wrap(err, "[HouseRepositoryImpl][GetAllHouses]")
	}
	return houses, errors.Wrap(err, "[HouseRepositoryImpl][GetAllHouses]")
}

func (h HouseRepositoryImpl) RemoveHouse(ctx *gin.Context, pid string) error {
	db := ctx.MustGet("DB").(*gorm.DB)
	var err error
	tx := db.Begin()

	// Delete entry from House table
	if err = tx.Where("pid = ?", pid).Delete(&entities.House{}).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[HouseRepositoryImpl][RemoveHouse]")
	}

	// Delete all entries from Picture table with foreign key houseID
	if err = tx.Where("house_pid = ?", pid).Delete(&entities.HousePicture{}).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "[HouseRepositoryImpl][RemoveHouse]")
	}

	// Commit transaction if both deletes succeed
	if err = tx.Commit().Error; err != nil {
		return errors.Wrap(err, "[HouseRepositoryImpl][RemoveHouse]")
	}

	return errors.Wrap(err, "[HouseRepositoryImpl][RemoveHouse]")
}

func (h HouseRepositoryImpl) CreatePictures(ctx *gin.Context, housePictures []entities.HousePicture) error {
	db := ctx.MustGet("DB").(*gorm.DB)

	err := db.Create(&housePictures).Error
	if err != nil {
		return errors.Wrap(err, "[HouseRepositoryImpl][CreatePictures]")
	}
	return errors.Wrap(err, "[HouseRepositoryImpl][CreatePictures]")
}
