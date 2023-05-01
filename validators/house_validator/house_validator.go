package house_validator

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/validators"
)

func ValidateCreateHouseRequest(ctx *gin.Context) (models.CreateHouseReq, error) {
	var createHouseReq models.CreateHouseReq

	err := ctx.ShouldBindJSON(&createHouseReq)
	if err != nil {
		return createHouseReq, errors.Wrap(err, "[ValidateCreateHouseRequest]")
	}

	return createHouseReq, errors.Wrap(err, "[ValidateCreateHouseRequest]")
}

func ValidateGetHouseRequest(ctx *gin.Context) (string, error) {
	housePid := ctx.Param("house_pid")
	if housePid == "" {
		return "", validators.ParamMissingOrInvalidError("housePid")
	}
	return housePid, nil
}

func ValidateGetFilteredHouseRequest(ctx *gin.Context) (models.HouseFilter, error) {
	var err error
	var houseFilter models.HouseFilter = models.HouseFilter{
		City:    ctx.Query("city"),
		State:   ctx.Query("state"),
		Country: ctx.Query("country"),
		Pincode: ctx.Query("pincode"),
	}

	if numBedrooms := ctx.Query("bedrooms"); numBedrooms != "" {
		houseFilter.Bedrooms, err = strconv.Atoi(numBedrooms)
		if err != nil {
			return houseFilter, validators.ParamMissingOrInvalidError("bedrooms")
		}
	}

	if numBathrooms := ctx.Query("bathrooms"); numBathrooms != "" {
		houseFilter.Bedrooms, err = strconv.Atoi(numBathrooms)
		if err != nil {
			return houseFilter, validators.ParamMissingOrInvalidError("bathrooms")
		}
	}

	if minPrice := ctx.Query("min_price"); minPrice != "" {
		houseFilter.MinPrice, err = strconv.Atoi(minPrice)
		if err != nil {
			return houseFilter, validators.ParamMissingOrInvalidError("min_price")
		}
	}

	if maxPrice := ctx.Query("min_price"); maxPrice != "" {
		houseFilter.MaxPrice, err = strconv.Atoi(maxPrice)
		if err != nil {
			return houseFilter, validators.ParamMissingOrInvalidError("max_price")
		}
	}

	if squareFootage := ctx.Query("square_footage"); squareFootage != "" {
		houseFilter.SquareFootage, err = strconv.Atoi(squareFootage)
		if err != nil {
			return houseFilter, validators.ParamMissingOrInvalidError("square_footage")
		}
	}

	fmt.Println(houseFilter.City)
	return houseFilter, errors.Wrap(err, "[ValidateGetFilteredHouseRequest]")
}
