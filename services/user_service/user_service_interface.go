package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context, userReq models.CreateUserReq) (entities.User, error)
	GetUserByPid(ctx *gin.Context, userPid string) (entities.User, error)
	AuthenticateUser(ctx *gin.Context, userLoginReq models.UserLoginReq)
}
