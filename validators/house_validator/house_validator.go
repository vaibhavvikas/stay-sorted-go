package house_validator

import (
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
