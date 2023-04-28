package user_repository

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vaibhavvikas/stay-sorted/constants"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
	"github.com/vaibhavvikas/stay-sorted/utils"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func (u UserRepositoryImpl) CreateUser(ctx *gin.Context, user entities.User) (entities.User, error) {
	db := ctx.MustGet("DB").(*gorm.DB)
	user.Pid = utils.GeneratePid(constants.Prefixes.USER, 25)

	err := db.Create(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "[UsersRepositoryImpl][CreateUser]")
	}
	return user, errors.Wrap(err, "[UsersRepositoryImpl][CreateUser]")
}

func (u UserRepositoryImpl) GetUserByPid(ctx *gin.Context, userPid string) (entities.User, error) {
	var user entities.User
	db := ctx.MustGet("DB").(*gorm.DB)

	err := db.Where("user_pid = ?", userPid).Take(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "[UsersRepositoryImpl][GetUserByPid]")
	}
	return user, errors.Wrap(err, "[UsersRepositoryImpl][GetUserByPid]")
}

func (u UserRepositoryImpl) GetUserByEmail(ctx *gin.Context, email string) (entities.User, error) {
	var user entities.User
	db := ctx.MustGet("DB").(*gorm.DB)

	err := db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, errors.Wrap(err, "[UsersRepositoryImpl][GetUserByEmail]")
	}
	return user, errors.Wrap(err, "[UsersRepositoryImpl][GetUserByEmail]")
}
